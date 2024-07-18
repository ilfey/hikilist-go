package baseModels

type ListModel[T any] struct {
	Results []*T `json:"results"`

	Count *int64 `json:"count,omitempty"`
}

type ListModelOption func(*ListModel[any])

func NewListModel[T any](items []*T) *ListModel[T] {
	model := &ListModel[T]{
		Results: items,
	}

	return model
}

func (m *ListModel[T]) WithCount(count int64) *ListModel[T] {
	m.Count = &count

	return m
}