package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Printf("someFunc with argument %v ran\n", num)
}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func main() {
	// Join time should be considered
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	go someFunc("4")

	time.Sleep(time.Second * 1)

	fmt.Println("hi")

	// Use channels to communicate between goroutines
	doneChannel := make(chan bool)
	apiChannel := make(chan bool)

	go func() {
		doneChannel <- true
	}()

	go func() {
		apiChannel <- true
	}()

	select {
	case msgFromDoneChannel := <-doneChannel:
		fmt.Println(msgFromDoneChannel)
	case msgFromApiChannel := <-apiChannel:
		fmt.Println(msgFromApiChannel)
	}

	// for select concurrency pattern
	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	// done channel concurrency pattern
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)

	// pipeline concurrency pattern
	// input
	nums := []int{2, 3, 4, 7, 1}

	//stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
