package utils

import (
	"strconv"
)

// StringToUint converts a string to a uint. Returns 0 if conversion fails.
func StringToUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(val)
}
