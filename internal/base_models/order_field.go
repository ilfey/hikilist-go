package baseModels

import "strings"

type OrderField string

func (o OrderField) IsDesc() bool {
	field := o.Field()
	if field == "" {
		return false
	}

	return strings.HasPrefix(string(o), "-")
}

func (o OrderField) IsAsc() bool {
	field := o.Field()
	if field == "" {
		return false
	}

	return !o.IsDesc()
}

func (o OrderField) Field() string {
	return strings.TrimPrefix(string(o), "-")
}

/*
ToGormQuery returns query string for GORM.

This string looks like: "{field_name} {ASC|DESC}"
*/
func (o OrderField) ToGormQuery() string {
	field := o.Field()
	if field == "" {
		return ""
	}

	if o.IsDesc() {
		return o.Field() + " DESC"
	}

	return o.Field() + " ASC"
}
