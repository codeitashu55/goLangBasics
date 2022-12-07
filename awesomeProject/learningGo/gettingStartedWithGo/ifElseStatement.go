package gettingStartedWithGo

import "fmt"

//in go it is better to return as early as possible and try to avoid un-necessary branches

//in blow code we avoided else branch and return as soon as we get result

func FindEvenOdd(no int) {
	//if statement
	if newNo := no; no%2 == 0 {
		fmt.Println(newNo, "no is even")
		return
	}
	fmt.Println(no, "no is odd")

}

//code to calculate grater of three no

func ToCheckGraterOfThreeNo(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}

	}
}

//code to depict if else lader -> to check if no is between smaller than 50 or btw 100 and 50 or grater than 100G
//the no defined in if assignment has its scope in the whole  else if lader and else if can also have there own assignment expression

func ToCheckIfNoIsInGivenRange(no int) {
	if findRange := no; findRange < 50 {
		println(findRange, "is between 0-49")
		return
	} else if newNo := 2; findRange >= 50 && findRange < 101 {
		println(findRange, "is between 50 - 101", newNo)
		return
	}

	//avoiding else branch if possible
	println(no, "is grater than 100")

}
