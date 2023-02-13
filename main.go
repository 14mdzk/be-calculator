package main

import (
	"fmt"

	calc "be-calculator/calculator"
)

func main() {
	var input string
	var result float64

	fmt.Println("Calculator")
	fmt.Println("Example: 2*+2*1/2")

	fmt.Print("Enter your function: ")
	fmt.Scanf("%s", &input)

	calculator := calc.Calculator{Functions: input}
	result = calculator.Calculate()

	fmt.Println("Result: ", result)
}
