package gettingStartedWithGo

import "fmt"

func ToPrintNoFrom1to10() {
	println()
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
}

// use of break statement

func ToPrintNoFrom1to5() {
	println()
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}
		fmt.Print(i, " ")
	}
}

//use of continue

func ToPrintOnlyOddNos() {
	println()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Print(i, " ")
	}
}

// nesting of loops to print an star pattern
// you can skip all three parts of for loop
//skiping last to part can make it work like an while loop

func ToPrintStarPattern() {
	println("\n")
	i := 0
	for i < 5 {
		for j := 0; j <= i; j++ {
			print("*")
		}
		i++
		println()
	}
}

//using tags to break outer loop

func ToShowOuterBreak() {

outer:
	for i := 1; i < 10; i++ {
		println()
		for j := 1; j <= i; j++ {
			if i == 6 {
				break outer
			}
			fmt.Print(j, " ")
		}
	}
}
