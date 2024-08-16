package action

import (
	"time"

	userModels "github.com/ilfey/hikilist-go/pkg/models/user"
)

type DetailModel struct {
	ID uint

	UserID uint
	User   *userModels.ListItemModel

	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}
