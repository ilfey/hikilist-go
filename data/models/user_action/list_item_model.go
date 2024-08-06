package userActionModels

type ListItemModel struct {
	ID uint `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt string `json:"created_at"`
}
