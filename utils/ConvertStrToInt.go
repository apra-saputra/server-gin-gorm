package utils

import (
	"strconv"
)

func ConverStrToInt(str string) (uint, error) {
	value, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}
