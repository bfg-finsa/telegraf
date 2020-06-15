package all

import (
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/basicstats"
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/final"
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/histogram"
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/merge"
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/minmax"
	_ "github.com/bfg-finsa/telegraf/plugins/aggregators/valuecounter"
)
