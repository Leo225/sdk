package sqlplus

import (
	"fmt"
	"strings"
)

func SQLPlaceholders(n int) string {
	var b strings.Builder

	for i := 0; i < n-1; i++ {
		b.WriteString("?,")
	}
	if n > 0 {
		b.WriteString("?")
	}
	return b.String()
}

func SQLSelectColumns(columns []string, as ...string) string {
	if len(columns) == 0 {
		return ""
	}

	stringBuilder := new(strings.Builder)
	tagLastIndex := len(columns) - 1
	for k, v := range columns {
		if len(as) > 0 && as[0] != "" {
			stringBuilder.WriteString(fmt.Sprintf("%s. `%s`", as[0], v))
		} else {
			stringBuilder.WriteString(fmt.Sprintf("`%s`", v))
		}

		if k != tagLastIndex {
			stringBuilder.WriteString(", ")
		}
	}
	return stringBuilder.String()
}
