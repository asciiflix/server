package utils

import "strconv"

//Parse string to uint (for user-id stuff)
func ParseStringToUint(toParse string) (uint, error) {
	data, err := strconv.ParseUint(toParse, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(data), nil
}
