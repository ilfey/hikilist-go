package sql

import (
	"fmt"
	"strings"
)

func where(table string, conds any) string {
	switch cond := conds.(type) {
	case int:
		return fmt.Sprintf(
			"%s.id = %d",
			table,
			cond,
		)
	case []int:
		expresion := ""

		for _, item := range cond {
			expresion += fmt.Sprintf("%d, ", item)
		}

		return fmt.Sprintf(
			"%s.id IN (%s)",
			table,
			strings.TrimSuffix(expresion, ", "),
		)
	case string:
		return cond
	case map[string]any:
		return whereMap(table, cond)
	}

	return ""
}

func whereMap(table string, conds map[string]any) string {
	var condition string

	for key, value := range conds {
		condition += fmt.Sprintf(
			"%s.%s = %s",
			table,
			toColumnName(key),
			toString(value),
		)
	}

	return condition
}
