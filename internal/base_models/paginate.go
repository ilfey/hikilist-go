package baseModels

import (
	"strconv"
)

type Paginate struct{}

func (Paginate) QueryInt(q map[string][]string, key string) int {
	valueStrings, ok := q[key]
	if !ok {
		return 0
	}

	valueString := valueStrings[0]

	value, err := strconv.Atoi(valueString)
	if err != nil {
		return 0
	}

	return value
}

func (Paginate) QueryOrder(q map[string][]string, key string) OrderField {
	valueStrings, ok := q[key]
	if !ok {
		return OrderField("")
	}

	return OrderField(valueStrings[0])
}

func (p *Paginate) GetOffset(page, limit int) int {
	return (page - 1) * limit
}
