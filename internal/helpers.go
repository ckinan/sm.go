package internal

import (
	"fmt"
	"strings"
)

func HumanBytes(kib int) string {
	switch {
	case kib >= 1<<20: // >= 1GiB
		return fmt.Sprintf("%.1f GiB", float64(kib)/float64(1<<20))
	case kib >= 10: // >= 1 MiB
		return fmt.Sprintf("%.1f MiB", float64(kib)/float64(1<<10))
	default:
		return fmt.Sprintf("%d KiB", kib)
	}
}

func extractFieldFromLine(line string) (string, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", fmt.Errorf("invalid line, expected at least 2 fields, got %v, line: %s", len(fields), line)
	}
	return fields[1], nil
}
