package main

import (
	"fmt"
	"unit_test/stringutil"
	"unit_test/calculator"
)

func main() {

	additionOp := calculator.GetOperation("addition")
	result := additionOp.Operate(5, 3)
	fmt.Println("5 + 3 =", result)

	multiplicationOp := calculator.GetOperation("multiplication")
	result = multiplicationOp.Operate(5, 3)
	fmt.Println("5 * 3 =", result)

	fmt.Println(stringutil.Reverse("hello"))
	fmt.Println(stringutil.ToUpper("hello"))
	fmt.Println(stringutil.IsEmpty("hello"))
	fmt.Println(stringutil.FirstRune("hello"))
	fmt.Println(stringutil.Concat("hello"))
}