package helpers

import (
	"errors"
	"strconv"
)

func ParseId(id string) (int64, error) {
	if id == "" {
		return 0.0, errors.New("user id cannot be empty")
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0.0, errors.New("user id must be a number")
	}
	return intId, nil
}
