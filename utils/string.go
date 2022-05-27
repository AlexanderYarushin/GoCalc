package utils

import (
	"strconv"
	"strings"
)

func ClearSpace(str *string) {
	for strings.Index(*str, " ") != -1 {
		*str = strings.Replace(*str, " ", "", -1)
	}
}

func IsNumber(str *string) bool {
	if _, err := strconv.ParseInt(*str, 10, 0); err == nil {
		return true
	}

	if _, err := strconv.ParseFloat(*str, 32); err == nil {
		return true
	}

	return false
}
