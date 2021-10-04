package validate

import (
	"errors"
	"regexp"
)

func VATID(vatID string) error {
	switch {
	case len(vatID) != 11:
		return errors.New("invalid length of vatID")
	case !regexp.MustCompile(`DE\d{9}`).MatchString(vatID):
		return errors.New("invalid format of vatID")
	default:
		return nil
	}
}
