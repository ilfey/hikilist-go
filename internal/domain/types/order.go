package types

import "strings"

type Order string

func (o Order) IsEmpty() bool {
	return strings.TrimPrefix(string(o), "-") == ""
}

func (o Order) IsDesc() bool {
	return !o.IsEmpty() && strings.HasPrefix(string(o), "-")
}

func (o Order) IsAsc() bool {
	return !o.IsEmpty() && !strings.HasPrefix(string(o), "-")
}

func (o Order) Field() string {
	return strings.TrimPrefix(string(o), "-")
}

/*
ToQuery returns query string.

This string looks like: "<field_name> <ASC|DESC>"
*/
func (o Order) ToQuery() string {
	field := o.Field()
	if field == "" {
		return ""
	}

	if o.IsDesc() {
		return o.Field() + " DESC"
	}

	return o.Field() + " ASC"
}
