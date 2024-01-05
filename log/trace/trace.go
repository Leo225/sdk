package trace

import (
	"fmt"
	"strconv"
)

const (
	defaultSpanID = "0"
)

func NextSpanID(spanID string) string {
	if spanID == "" {
		return defaultSpanID
	}

	l := len(spanID)
	lastID, _ := strconv.Atoi(spanID[l-1:])
	if lastID > 0 {
		spanID = spanID[:l-2]
	}
	lastID++
	return fmt.Sprintf("%s.%d", spanID, lastID)
}

func StartSpanID(spanID string) string {
	if spanID == "" {
		return defaultSpanID
	}

	l := len(spanID)
	lastID, _ := strconv.Atoi(spanID[l-1:])
	lastID++
	return fmt.Sprintf("%s.%d", spanID, lastID)
}
