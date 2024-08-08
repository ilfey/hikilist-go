package collection

type ListItemModel struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id"`

	Title string `json:"title"`

	Description *string `json:"description"`
}
