package collectionModels

type ListItemModel struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id"`

	Title string `json:"title"`

	Description *string `json:"description"`

	IsPublic bool `json:"is_public"`
}

func (ListItemModel) TableName() string {
	return "collections"
}
