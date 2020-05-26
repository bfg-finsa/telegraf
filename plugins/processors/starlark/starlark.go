package starlark

import (
	"errors"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal/choice"
	"github.com/influxdata/telegraf/plugins/processors"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
)

const (
	description  = "Process metrics using a Starlark script"
	sampleConfig = `
  source = ""
  on_error = "pass"
`
	defaultOnError = "pass"
)

type Starlark struct {
	Source  string `toml:"source"`
	OnError string `toml:"on_error"`

	Log telegraf.Logger `toml:"-"`

	thread    *starlark.Thread
	applyFunc *starlark.Function
	args      starlark.Tuple
	results   []telegraf.Metric
}

func (s *Starlark) Init() error {
	err := choice.Check(s.OnError, []string{"pass", "drop"})
	if err != nil {
		return err
	}

	s.thread = &starlark.Thread{
		Name:  "processor.starlark",
		Print: func(_ *starlark.Thread, msg string) { s.Log.Debug(msg) },
	}

	predeclared := starlark.StringDict{}

	_, program, err := starlark.SourceProgram("processor.starlark", s.Source, predeclared.Has)
	if err != nil {
		return err
	}

	// Execute source
	globals, err := program.Init(s.thread, predeclared)
	if err != nil {
		return err
	}

	// Freeze the global state.  This prevents modifications to the processor
	// state and prevents scripts from containing errors storing tracking
	// metrics.  Tasks that require global state will not be possible due to
	// this, so maybe we should relax this in the future.
	globals.Freeze()

	// The source should define an apply function.
	apply := globals["apply"]

	if apply == nil {
		return errors.New("apply is not defined")
	}

	var ok bool
	if s.applyFunc, ok = apply.(*starlark.Function); !ok {
		return errors.New("apply is not a function")
	}

	if s.applyFunc.NumParams() != 1 {
		return errors.New("apply function must take one parameter")
	}

	// Reusing the same metric wrapper to skip an allocation.  This will cause
	// any saved references point to the new metric, but due to freezing the
	// globals none should exist.
	met := &Metric{}
	s.args = make(starlark.Tuple, 1)
	s.args[0] = met

	// Allocate a slice for return values.  FIXME can we remove?
	s.results = make([]telegraf.Metric, 0, 10)

	return nil
}

func (s *Starlark) SampleConfig() string {
	return sampleConfig
}

func (s *Starlark) Description() string {
	return description
}

func (s *Starlark) Apply(metrics ...telegraf.Metric) []telegraf.Metric {
	s.results = s.results[:0]
	for _, m := range metrics {
		s.args[0].(*Metric).metric = m

		rv, err := starlark.Call(s.thread, s.applyFunc, s.args, nil)
		if err != nil {
			if err, ok := err.(*starlark.EvalError); ok {
				for _, line := range strings.Split(err.Backtrace(), "\n") {
					s.Log.Error(line)
				}
			} else {
				s.Log.Error(err)
			}
			continue
		}

		switch rv := rv.(type) {
		case *starlark.List:
			iter := rv.Iterate()
			defer iter.Done()
			var v starlark.Value
			for iter.Next(&v) {
				switch v := v.(type) {
				case *Metric:
					m := v.Unwrap()
					if containsMetric(s.results, m) {
						s.Log.Errorf("Duplicate metric reference detected")
						continue
					}
					s.results = append(s.results, m)
				default:
					s.Log.Errorf("Invalid type returned in list: %T", v)
				}
			}
		case *Metric:
			s.results = append(s.results, rv.Unwrap())
		case starlark.NoneType:
			return nil
		default:
			s.Log.Errorf("Invalid type returned: %T", rv)
		}
	}
	return s.results
}

func containsMetric(metrics []telegraf.Metric, metric telegraf.Metric) bool {
	for _, m := range metrics {
		if m == metric {
			return true
		}
	}
	return false
}

func init() {
	// https://github.com/bazelbuild/starlark/issues/20
	resolve.AllowFloat = true
	resolve.AllowLambda = true
	resolve.AllowSet = true

	// FIXME Can we do this per program?
	starlark.Universe["Metric"] = starlark.NewBuiltin("Metric", newMetric)
	starlark.Universe["deepcopy"] = starlark.NewBuiltin("deepcopy", deepcopy)
}

func init() {
	processors.Add("starlark", func() telegraf.Processor {
		return &Starlark{
			OnError: defaultOnError,
		}
	})
}