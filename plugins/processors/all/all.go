package all

import (
	_ "github.com/bfg-finsa/telegraf/plugins/processors/clone"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/converter"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/date"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/dedup"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/defaults"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/enum"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/filepath"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/override"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/parser"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/pivot"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/port_name"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/printer"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/regex"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/rename"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/s2geo"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/strings"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/tag_limit"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/template"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/topk"
	_ "github.com/bfg-finsa/telegraf/plugins/processors/unpivot"
)
