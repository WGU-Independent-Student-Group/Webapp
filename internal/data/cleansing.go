package data

import (
	"fmt"
	"strconv"
)

// file to store all the data verification yuck

func parseIntField(s string, fieldName string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", fieldName, s, err)
	}
	return i, nil
}

func parseFloatField(s string, fieldName string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", fieldName, s, err)
	}
	return f, nil
}
