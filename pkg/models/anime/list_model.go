package anime

type ListModel struct {
	Results []*ListItemModel `json:"results"`

	Count *uint `json:"count,omitempty"`
}
