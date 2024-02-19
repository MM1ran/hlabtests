package main

import "errors"

func Divide(a, b int) (result int, err error) {
	if a == 0 && b == 0 {
		return 0, errors.New("both numbers are zero")
	} else {
		return a / b, nil
	}
}
