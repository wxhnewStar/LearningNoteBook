package main

import (
	"errors"
	"fmt"
)

var errDivisionByZero = errors.New("division by zero")

func div(divided, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}

	return divided / divisor, nil
}

func main() {
	fmt.Println(div(1, 0))

	_, err := div(1, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
}
