package main

import (
	"errors"
	"fmt"
)

var ErrDividebyZero = errors.New("can't divide by zero")

func main() {
	fmt.Println("Divid 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Divide 2 by 0")
	fmt.Println(divide(2, 0))
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDividebyZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
