package errorHandlingInGo

import (
	"fmt"
	"sync"
)

//defer in go

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf(" %s %s \n", p.firstName, p.lastName)
}

func DeferFunctionCallInGo() {
	p := person{
		"bob",
		"bob",
	}
	fmt.Print("Welcome user")
	defer p.fullName()

}

// finding area of rectangles using waitGroup
type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.width <= 0 || r.length <= 0 {
		println("one of the parameter fails")
		return
	}
	println("area of rectangle is", r.width*r.length)
}

func CalculateAreaOfRectangle() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9} // they are executed in lifo order
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1) // to add to the counter of waitGroup
		go v.area(&wg)
	}
	wg.Wait() // <- control returns here when the counter becomes zero
	fmt.Println("All go routines finished executing")
}
