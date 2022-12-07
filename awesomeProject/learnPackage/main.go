package main

import (
	"learnPackage/simpleIntrest"
)

func main() {
	sum := simpleIntrest.AddingTwoNo(12, 12)
	interest := simpleIntrest.Calculate(8000, 9, 80)
	println(sum, interest)

}
