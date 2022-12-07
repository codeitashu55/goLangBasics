package collectionInGo

import (
	"fmt"
)

func MultipleWayOfInitializing() {
	//extended way
	println()
	//short hand are not that usefull in these cases because we will not use them that often
	arr1 := [...]int{}
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i, _ := range arr1 {
		fmt.Print(i, " ")
	}
	println()

	/** few points

	-> (...) length to be calculated at run time
	-> a := [10]int{} -> short hand for array of length 10 and element can be provided after initialization
	-> var a [10]int  if no value is given the elements are set to 0
	-> [x]int and [y]int are different type if x != y

	**/

	//arrays are pass by value in function and when we assign them to an variable

	a := [5]int{}
	for i, _ := range a {
		a[i] = i + 1
	}

	b := a

	for i, _ := range b {
		b[i] = a[i] + 1
	}

	//the values of a don't get changed when we change the value of a
	printArray(a)
	printArray(b)

	modifyArray(a)
	fmt.Print("printing array after modify function is called\n")
	printArray(a)

}

func modifyArray(a [5]int) {
	for i, _ := range a {
		a[i] = 0
	}
	fmt.Print("printing array inside the function\n")
	printArray(a)
}

func printArray(a [5]int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	println()
}

//creating and printing multidimensional

func CreateAndPrintMultiDimensionalArray(rows, coloum int) {
	//creating 2d array using append function
	array2d := make([][]int, 0)
	for i := 0; i < rows; i++ {
		temp := make([]int, 0)
		for j := 0; j < coloum; j++ {
			temp = append(temp, i+j+1)
		}
		array2d = append(array2d, temp)
	}

	for _, v1 := range array2d {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		println()
	}

}

//slices in go

func SlicesInGo() {
	//a reference to array so the problem with array is that we can not make them dynamic
	//if we want to make the array of given length we have to use slice
	total := 7
	arr := make([]int, 0)
	for i := 0; i < total; i++ {
		arr = append(arr, i+1)
	}

	for _, v := range arr {
		fmt.Print(v, " ")
	}

	//any change in slice is reflected on the array as well
	arr1 := [...]int{1, 2, 3, 4, 5}
	sliceArr1 := arr1[1:3]

	//do changes in slice
	fmt.Print("\nprevious array\n")
	printArray(arr1)
	//changeSlice(sliceArr1)
	fmt.Print("array after changes in slice\n")
	printArray(arr1)

	println()
	for _, v := range sliceArr1 {
		println(v, " ")
	}

	//re-slicing capacity and length
	println("capacity of slice ->", cap(sliceArr1), "length of an slice ->", len(sliceArr1))

	//if we are doing re-slicing we are doing slicing of the underlying array
	//the start index that we choose is based upon the slice or array you are choosing from

	println()
	resliceArr1 := sliceArr1[0:3]
	for _, v := range resliceArr1 {
		fmt.Print(v, " ")
	}

}

func changeSlice(slice []int) {
	for i, _ := range slice {
		slice[i]++
	}
}

func AppendFunctionAndCapacityOfSlices() {

	//each and every slice has an underlying array working inside of it
	println()
	// capacity defiles the max length of slice that can be created form the given slice
	slice := make([]int, 3, 6)
	reSlice := slice[:cap(slice)]
	for _, v := range reSlice {
		fmt.Print(v, " ")
	}

	//append is an variadic function -> you can append as many element as you want
	slice = append(slice, 1, 2, 3, 4)

	println()
	for _, v := range slice {
		fmt.Print(v, " ")
	}

}

//slice and variadic function and variadic variable

func VariadicFunctionAndRelationWithSlice(a int, b ...int) {
	//function to check if a exists in b
	var found = false
	for i, v := range b {
		if v == a {
			fmt.Print("first found at index", i)
			found = true
			break
		}
	}

	if !found {
		fmt.Print(a, " is not present in b")
	}
	println()
}

//creating same function using slice

func SamethingUsingSlice(a int, b []int) {
	var found = false
	for i, v := range b {
		if v == a {
			fmt.Print("first found at index ", i)
			found = true
			break
		}
	}

	if !found {
		fmt.Print(a, " is not present in b ")
	}
	println()
}

func Change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
