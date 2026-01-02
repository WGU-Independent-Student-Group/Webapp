package data

import (
	"fmt"
	"strconv"
	"strings"
)

// file to store all the data verification yuck

var states = map[string]struct{}{
	// map lookup is preferable over list
	// empty structs do look weird tho
	"alabama":        {},
	"alaska":         {},
	"arizona":        {},
	"arkansas":       {},
	"california":     {},
	"colorado":       {},
	"connecticut":    {},
	"delaware":       {},
	"florida":        {},
	"georgia":        {},
	"hawaii":         {},
	"idaho":          {},
	"illinois":       {},
	"indiana":        {},
	"iowa":           {},
	"kansas":         {},
	"kentucky":       {},
	"louisiana":      {},
	"maine":          {},
	"maryland":       {},
	"massachusetts":  {},
	"michigan":       {},
	"minnesota":      {},
	"mississippi":    {},
	"missouri":       {},
	"montana":        {},
	"nebraska":       {},
	"nevada":         {},
	"new hampshire":  {},
	"new jersey":     {},
	"new mexico":     {},
	"new york":       {},
	"north carolina": {},
	"north dakota":   {},
	"ohio":           {},
	"oklahoma":       {},
	"oregon":         {},
	"pennsylvania":   {},
	"rhode island":   {},
	"south carolina": {},
	"south dakota":   {},
	"tennessee":      {},
	"texas":          {},
	"utah":           {},
	"vermont":        {},
	"virginia":       {},
	"washington":     {},
	"west virginia":  {},
	"wisconsin":      {},
	"wyoming":        {},
}

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

// convert lon and lat from string to float64
// then check for coordinate validity
func validateCoordinates(lonStr, latStr string) (float64, float64, error) {
	lat, err := parseFloatField(latStr, "latitude")
	if err != nil {
		return 0, 0, err
	}

	lon, err := parseFloatField(lonStr, "longitude")
	if err != nil {
		return 0, 0, err
	}

	if lon < -180 || lon > 180 {
		return 0, 0, err
	}

	if lat < -90 || lat > 90 {
		return 0, 0, err
	}

	return lon, lat, nil
}

// keep usa entries only
func validateCountry(country string) bool {
	trimmed := strings.TrimSpace(country)
	return strings.EqualFold(trimmed, "united states")
}

func validateState(state string) bool {
	_, valid := states[strings.ToLower(strings.TrimSpace(state))]
	return valid
}
