package statistic

type RW interface {
	Readable
	Writeable
}

type Readable interface {
	String() string

	Pages() uint
	Fetched() uint
	Resolved() uint
	Scipped() uint
}

type Writeable interface {
	AddPages(count uint)

	AddFetched(count uint)

	AddResolved(count uint)

	AddScipped(count uint)
}
