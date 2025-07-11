package main

import (
	"errors"
	"fmt"
)

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}
func main() {
	res, err := devide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
