//go:build !solution

package testequal

import (
	"strings"
)

func cmp(a,b interface{}) bool {
	switch valA := a.(type) {

	case int, int8, int16, int32, int64:

		switch valB := b.(type) {

		case int, int8, int16, int32, int64:
			if valA == valB {
				return true
			}
		default:
			return false
		}
	case uint, uint8, uint16, uint32, uint64:

		switch valB := b.(type) {
		case uint, uint8, uint16, uint32, uint64:
			if valA == valB {
				return true
			}
		default:
			return false
		}
	case string:
		valB, ok := b.(string);

		if !ok{
			return false
		}

		if strings.Compare(valB,valA) == 0 {
			return true
		}
		return false

	case map[string]string:
		valB, ok := b.(map[string]string);

		if !ok{
			return false
		}

		if valB == nil || valA == nil{
			return false
		}

		if len(valB) != len(valA) {
			return false
		}

		for k, v := range valA {
			if i, ok := valB[k]; !(ok && i == v)  {
				return false
			}
		}
		return true
	case []int:
		valB, ok := b.([]int);

		if !ok{
			return false
		}

		if valB == nil || valA == nil{
			return false
		}

		if len(valA) == len(valB) {
			for i, v := range valA {
				if valB[i] != v {
					return false
				}
			}
		} else {
			return false
		}
		return true
	case []byte:
		valB, ok := b.([]byte);

		if !ok{
			return false
		}

		if valB == nil || valA == nil{
			return false
		}

		if len(valA) == len(valB) {
			for i, v := range valA {
				if valB[i] != v {
					return false
				}
			}
		} else {
			return false
		}
		return true
	}
	return false
}
// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !cmp(expected, actual) {
		for _, i := range msgAndArgs {
			if len(msgAndArgs) > 1 {
				t.Errorf(i.(string), msgAndArgs[1:]...)
				return false
			} else {
				t.Errorf(i.(string))
				return false
			}
		}
		t.Errorf("")
		return false
	}
	return true
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if cmp(expected, actual) {
		for _, i := range msgAndArgs {
			if len(msgAndArgs) > 1 {
				t.Errorf(i.(string), msgAndArgs[1:]...)
				return false
			} else {
				t.Errorf(i.(string))
				return false
			}
		}
		t.Errorf("")
		return false
	}
	return true
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !cmp(expected, actual) {
		for _, i := range msgAndArgs {
			if len(msgAndArgs) > 1 {
				t.Errorf(i.(string), msgAndArgs[1:]...)
				return
			} else {
				t.Errorf(i.(string))
				return
			}
		}
		t.Errorf("")
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if cmp(expected, actual) {
		for _, i := range msgAndArgs {
			if len(msgAndArgs) > 1 {
				t.Errorf(i.(string), msgAndArgs[1:]...)
				return
			} else {
				t.Errorf(i.(string))
				return
			}
		}
		t.Errorf("")
	}
}
