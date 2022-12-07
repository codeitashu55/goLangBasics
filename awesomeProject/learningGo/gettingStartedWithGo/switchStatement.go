package gettingStartedWithGo

import (
	"fmt"
)

//creating basic switch condition

/***
few more things
-> multiple cases can run same code
  'i','o','e','u','a' -> can be grouped to execute same code

-> case can not repeat itself
***/

func FindingTheFinger() {
	println()
	finger := 7
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Ye kon si finger howae hai")

	}

}

func ConditionLessSwitches(no int) {
	switch {
	case (no > 0) && (no < 51):
		println(no, "is between 1 - 50")

	case no > 50 && no <= 100:
		println(no, "is between 50 - 100")

	default:
		println("no is larger than 100")

	}
}

//use of fallthrough in switch (fallthrough will happen ir-respective if the condition is evaluated to true or false)

func FallthroughInSwitch(no int) {
	switch {
	case no < 50:
		println(no, "is less than 50")
		fallthrough

	case no > 100:
		println(no, "is grater than 100")

	}

	/***
	few important things

	-> if code reaches break the control goes out of the switch statement

	***/
}

