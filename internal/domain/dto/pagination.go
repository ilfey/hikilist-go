package dto

import "github.com/ilfey/hikilist-go/internal/domain/types"

type PaginationRequestDTO struct {
	Page  uint64      `json:"page"`
	Limit uint64      `json:"limit"`
	Order types.Order `json:"order"`
}
