package paginate

var DefaultOptions = Options{
	DefaultOrder: "-id",
	DefaultLimit: 10,
	MaxLimit:     100,
	AwaiableOrders: []string{
		"",
		"id",
		"-id",
	},
}

type Options struct {
	DefaultOrder   string
	DefaultLimit   int
	MaxLimit       int64
	AwaiableOrders []string
}

type Option func(*Options)

func WithDefaultOrder(order string) Option {
	return func(options *Options) {
		options.DefaultOrder = order
	}
}

func WithDefaultLimit(limit int) Option {
	return func(options *Options) {
		options.DefaultLimit = limit
	}
}

func WithMaxLimit(limit int64) Option {
	return func(options *Options) {
		options.MaxLimit = limit
	}
}

func WithAwaiableOrders(orders []string) Option {
	return func(options *Options) {
		options.AwaiableOrders = orders
	}
}
