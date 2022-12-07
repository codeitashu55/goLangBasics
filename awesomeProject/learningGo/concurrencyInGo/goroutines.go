package concurrencyInGo

import (
	"fmt"
	"sync"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

//channels in go

func ChannelExample() {
	var channel chan int
	if channel == nil {
		fmt.Println("channel is nil")
		channel = make(chan int)
		fmt.Printf("Type of a is %T \n", channel)
	}
}

func HelloWorld() {
	a := make(chan bool)
	go launchGoRoutine(a) //it launches the function in an go-routine
	received := <-a
	fmt.Println("end of the hello world function", received)
	//go numbers()
	//go alphabets()
	//time.Sleep(3000 * time.Millisecond)

}

func launchGoRoutine(a chan bool) {
	fmt.Println("hello world")
	a <- true
	//the control is returned when a is receives value
	time.Sleep(1 * time.Second)
	fmt.Println("returning controll")
}

//demo of an api call

//for result := <-api for this to exist there must exist an go routine which writes into this channel

func ApiCall() {
	api1 := make(chan int)
	api2 := make(chan int)
	//this will work but will take more time
	go doNetworkCall(api1)
	go doNetworkCall(api2)
	result := <-api1
	oneMore := <-api2

	fmt.Printf("result of api1 is %d \n", result)
	fmt.Printf("result of api2 is %d \n", oneMore)
}

func doNetworkCall(api chan<- int) {
	fmt.Println("calling api..")
	time.Sleep(2 * time.Second)
	api <- 200
}

func printValues(ch chan<- int) {
	for i := 1; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func ClosingChannel() {
	ch := make(chan int)
	go printValues(ch)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Received", v, ok)
			break
		}

		fmt.Println("Received", v, ok)
	}
}

func toEmitDigits(dig int, ch chan<- int) {
	for dig != 0 {
		no := dig % 10
		ch <- no
		dig = dig / 10
	}

	close(ch)
}

func returnSqOfDigits(ch chan<- int, no int) {
	chl := make(chan int)
	go toEmitDigits(no, chl)
	sum := 0
	for digit := range chl {
		sum += digit * digit
	}
	ch <- sum
}

func returnCuOfDigits(ch chan<- int, no int) {
	chl := make(chan int)
	go toEmitDigits(no, chl)
	sum := 0
	for digit := range chl {
		sum += digit * digit * digit
	}
	ch <- sum
}

func ToCalSumOfSqAndCu() {
	sqCh := make(chan int)
	cuCh := make(chan int)

	go returnSqOfDigits(sqCh, 589)
	go returnCuOfDigits(cuCh, 589)

	square := <-sqCh
	cube := <-cuCh

	fmt.Println("The sum of cube and square is", square+cube)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func UsingWaitGroupToWaitForTheGoRoutines() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
