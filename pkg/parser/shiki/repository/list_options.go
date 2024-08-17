package repository

import (
	"strconv"
)

type Kind string
type Order string
type Status string

const (
	KindTV        Kind = "tv"
	KindMovie     Kind = "movie"
	KindOVA       Kind = "ova"
	KindONA       Kind = "ona"
	KindSpecial   Kind = "special"
	KindTVSpecial Kind = "tv_special"
	KindMusic     Kind = "music"
	KindPV        Kind = "pv"
	KindCM        Kind = "cm"
	KindTV13      Kind = "tv_13"
	KindTV24      Kind = "tv_24"
	KindTV48      Kind = "tv_48"
	KindNone      Kind = ""

	OrderID         Order = "id"
	OrderRanked     Order = "ranked"
	OrderKind       Order = "kind"
	OrderPopularity Order = "popularity"
	OrderName       Order = "name"
	OrderAiredOn    Order = "aired_on"
	OrderEpisodes   Order = "episodes"
	OrderStatus     Order = "status"
	OrderRandom     Order = "random"
	OrderNone       Order = ""

	StatusAnons    Status = "anons"
	StatusOngoing  Status = "ongoing"
	StatusReleased Status = "released"
	StatusNone     Status = ""
)

type ListOptions struct {
	Page     uint
	Limit    uint
	Order    Order
	Kind     Kind
	Status   Status
	Censored bool
}

func (o *ListOptions) ToMap() map[string]string {
	values := map[string]string{
		"page":     strconv.FormatUint(uint64(o.Page), 10),
		"limit":    strconv.FormatUint(uint64(o.Limit), 10),
		"order":    string(o.Order),
		"kind":     string(o.Kind),
		"status":   string(o.Status),
		"censored": strconv.FormatBool(o.Censored),
	}

	for key, value := range values {
		if value == "" {
			delete(values, key)
		}
	}

	return values
}

type ListOption = func(*ListOptions)

// Страница
func PageOption(page uint) ListOption {
	return func(opt *ListOptions) {
		opt.Page = page
	}
}

func LimitOption(limit uint) ListOption {
	return func(opt *ListOptions) {
		opt.Limit = limit
	}
}

func OrderOption(order Order) ListOption {
	return func(opt *ListOptions) {
		opt.Order = order
	}
}

func KindOption(kind Kind) ListOption {
	return func(opt *ListOptions) {
		opt.Kind = kind
	}
}

func StatusOption(status Status) ListOption {
	return func(opt *ListOptions) {
		opt.Status = status
	}
}

func CencoredOption(cencored bool) ListOption {
	return func(opt *ListOptions) {
		opt.Censored = cencored
	}
}
