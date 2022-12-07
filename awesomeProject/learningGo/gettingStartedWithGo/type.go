package gettingStartedWithGo

import (
	"fmt"
	"strconv"
	_ "strconv"
	"unsafe"
)

func init() {
	println("calling type")
}
func Type() {
	//boolean
	a := false
	b := true

	println("a-> ", a, "b->", b)

	// && and operator  || or operator
	c := a && b
	d := a || b

	println("c-> ", c, "d->", d)

	//signed integer
	/*
		-(2^7) to 2^7 -1 range for int8
		in the same way range for int16 , int32 can be calculated

		simple int -> represent 32 or 64 bit singed int depending upon the system

	*/

	// %T for type and %d to print an integer
	// unsafe.Sizeof(x) -> gives size in bytes  1 byte = 8 bits
	//Printf() is used to format output %T to get the type
	var x int
	fmt.Printf("type of variable - %T  size of variable - %d\n", x, unsafe.Sizeof(x))
	/*
	  un-signed integer

	  uint8 -> 0 to 2^8 -1 => 0 to 255
	  and similarly we can calculate for others as well
	*/

	var y uint8 = 255 //the max value it can store
	println(y)

	//floating point type (by default it is float64)
	//range is based on floating points -> accurate upto 6 decimal places for 64
	x1, x2 := 4.22, 5.28
	fmt.Printf("type of a %T type of b %T", x1, x2)

	sum := x1 + x2
	diff := x1 - x2

	println("sum ", sum, "diff ", diff)

	y1, y2 := 55, 70
	println("sum ", y1+y2, "diff", y2-y1)

	//complex no the real part and imaginary part must have same type [int / float]
	c1 := 4 + 3i
	var c2 = complex(1, 3)
	println("sum", c1+c2, "product", c1*c2)

	//byte and rune are alias of uint8 and int32 (TODO -> Pending)
	//string in collection of bytes in go -> explanation (TODO -> Pending)

	//sting

	var first = "Asutosh"
	last := "Pandey"

	// + can be used to connect two strings
	name := first + "-" + last
	println(name)

	//Type conversion
	//in go there is no automatic type conversion

	/*
			no two variable in an expression can be of different type
		    promotion is doned in case an constant of higher bit is present
	*/

	y3 := x2 + x1 + 2.0

	//an expression assigned to an variable must return the value of thats variable
	var typeCast int = int(x1)

	println(y3)
	println(typeCast)
	constants()
}

func constants() {
	//declaring constants in go
	const x = 44

	//declaring multiple constants
	const y, z = 20, "New Constant"
	println(x, y, z)

	/*
	  in constant re-assignment is not allowed the value must be provided at compile time
	  value returned by a function can't be assigned to a constant
	*/

	//const m = math.Sqrt(36)
	//println(m)

	//string constant , Typed and Untyped Constants

	//example of an untyped constant
	const n = "Hello"
	name := n
	fmt.Printf("Type of name %T and value is %v", name, name)

	//example of typed constant
	const m string = "World"

	//example why it is strictly typed
	var defaultName string = "Ashu"
	type myString string //that how we crate an alias of string
	var customName myString = "Abhi"

	println(defaultName, customName)

	//these bottom two are not allowed
	//even alias of same type can not be assigned

	//defaultName = customName
	//customName = defaultName

	const a = "5"
	//var zero = "5"
	//var intVar int = a
	//var int32Var int32 = a
	//var float64Var float64 = a
	//var complex64Var complex64 = a

	test, _ := strconv.Atoi(a)
	println(test, "ho gaya hua")

	//hence the constant is un-typed it reforms it value according to the need of the variable
	//if a constant is un-typed it can adjust it's value according to expression as well

	//println(intVar, int32Var, float64Var, complex64Var)
	//println(8.0 / a) //the type get converted into the type of the larger bit var

}
