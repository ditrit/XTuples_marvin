package convert

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to int: %v", s, err)
	}
	return i, nil
}

func StringToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("could not convert %s to bool: %v", s, err)
	}
	return b, nil
}
