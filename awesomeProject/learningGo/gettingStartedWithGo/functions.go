package gettingStartedWithGo

func init() {
	println("Calling function")
}

func Function() {
	functionsInGo()
}

func functionsInGo() {
	println("Total amount", calculatePrice(10, 10))

	//go allows multiple return type
	area, parameter := calculateParameterAndArea(5, 6)
	//use of blank identifier you can discard certain returned value using this avoiding unnecessary variable
	area, _ = calculateParameterAndArea(2, 3)
	println(area, parameter)
}

func calculatePrice(price, count int) int {
	total := price * count
	return total
}

//if a function has multiple return type the must be puted inside ()
//func calculateParameterAndArea(length, width float64) (float64, float64) {
//	parameter := 2 * (length + width)
//	area := length * width
//
//	return area, parameter
//}

//we can use named parameter in function
func calculateParameterAndArea(length, width float64) (area, parameter float64) {
	parameter = 2 * (length + width)
	area = length * width
	return //the values are returned as soon as the control reaches return
}
