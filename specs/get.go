package specs

import "github.com/lab210-dev/structkit"

type Get func(source any, field string, opt ...*structkit.Option) any
