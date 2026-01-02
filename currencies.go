package currencies

import (
	"strconv"
	"strings"
)

type CurrencyOptions struct {
	ThousandsSeparator string
	DecimalSeparator   string
	Code               string
}

func Format(value int, options CurrencyOptions) string {
	// Defaults
	if options.ThousandsSeparator == "" {
		options.ThousandsSeparator = "."
	}
	if options.DecimalSeparator == "" {
		options.DecimalSeparator = ","
	}

	// Convert int to string
	raw := strconv.Itoa(value)

	// Add thousand separators
	var parts []string
	for len(raw) > 3 {
		parts = append([]string{raw[len(raw)-3:]}, parts...)
		raw = raw[:len(raw)-3]
	}
	if raw != "" {
		parts = append([]string{raw}, parts...)
	}

	formatted := strings.Join(parts, options.ThousandsSeparator)

	result := "$" + formatted

	if options.Code != "" {
		result += " " + options.Code
	}

	return result
}
