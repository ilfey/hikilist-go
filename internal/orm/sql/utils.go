package sql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ilfey/hikilist-go/internal/inflection"
)

func toColumnName(name string) string {
	if name == "ID" {
		return "id"
	}

	if strings.HasSuffix(name, "ID") {
		return inflection.Snake(strings.TrimSuffix(name, "ID")) + "_id"
	}

	return inflection.Snake(name)
}

func toString(value any) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case time.Time:
		return fmt.Sprintf("'%s'", v.UTC().Format(time.RFC3339))
	}

	return ""
}
