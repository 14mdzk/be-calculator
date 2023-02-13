package calculator

import (
	"fmt"
	"regexp"
	"strconv"
)

type Calculator struct {
	Functions string
	numbers   []float64
	operators []string
}

func (c *Calculator) Calculate() (result float64) {
	c.transform()
	result = 0.2
	result = c.calculate()
	return
}

func (c *Calculator) transform() {
	rxOperator := regexp.MustCompile(`(\*|\+|\-|\/)`)
	c.operators = rxOperator.FindAllString(c.Functions, -1)

	rxNumber := regexp.MustCompile(`([0-9]*[.])?[0-9]+`)
	numbers := rxNumber.FindAllString(c.Functions, -1)

	for _, number := range numbers {
		if v, err := strconv.ParseFloat(number, 64); err == nil {
			if err != nil {
				continue
			}

			c.numbers = append(c.numbers, v)
		}
	}
}

func (c *Calculator) calculate() (result float64) {
	fmt.Println(c.numbers)
	for key, number := range c.numbers {
		if key == 0 {
			result = number
			continue
		}

		switch c.operators[key-1] {
		case "+":
			result = c.add(result, number)
		case "-":
			result = c.subtract(result, number)
		case "*":
			result = c.multiply(result, number)
		case "/":
			result = c.divide(result, number)
		default:
			result = 0
		}
	}

	return
}

func (c *Calculator) add(a, b float64) (result float64) {
	result = a + b
	return
}
func (c *Calculator) subtract(a, b float64) (result float64) {
	result = a - b
	return
}
func (c *Calculator) multiply(a, b float64) (result float64) {
	result = a * b
	return
}
func (c *Calculator) divide(a, b float64) (result float64) {
	result = a / b
	return
}
