package internal

import (
	"fmt"
	"strings"
)

func extractFieldFromLine(line string) (string, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", fmt.Errorf("invalid line, expected at least 2 fields, got %v, line: %s", len(fields), line)
	}
	return fields[1], nil
}
