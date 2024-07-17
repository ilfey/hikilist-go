package parser

type Query[T any] func() (T, error)

func RunQueryAsync[T any](q Query[T]) (T, error) {
	
	return q()
}
