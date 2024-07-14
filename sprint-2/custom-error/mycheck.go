//go:build !solution

package mycheck

import (
	"strings"
	"unicode"
	"errors"
)

var ErrFoundNum = errors.New("found numbers")
var ErrLargeLine = errors.New("line is too long")
var ErrNoTwoSpace = errors.New("no two spaces")

type sErrors []error

func (se sErrors) Error() string {
	if len(se) == 0 {
		return "Unknown error"
	}
	str := ""
	for _, i := range se {
		str += i.Error() + ";"
	}
	return str[:len(str)-1]
}


func MyCheck(input string) error {
	var errs sErrors


	for _, i := range input {
		if unicode.IsDigit(i) {
			errs = append(errs, ErrFoundNum)
			break
		}
	}

	if len(input) >= 20 {
		errs = append(errs, ErrLargeLine)
	}

	if strings.Count(input, " ") != 2 {
		errs = append(errs, ErrNoTwoSpace)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
