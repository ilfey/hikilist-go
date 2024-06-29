package book

// Book create model
type CreateModel struct {
	Name string `json:"name" form:"name"`
}