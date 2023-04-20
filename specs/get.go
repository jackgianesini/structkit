package specs

import "github.com/kitstack/structkit"

type Get func(source any, field string, opt ...*structkit.Option) any
