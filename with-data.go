package main

import (
	"fmt"
	"sync"
)

type responseProcess1 struct {
	Data  int
	Error error
}

type responseProcess2 struct {
	Data  string
	Error error
}

func processData1(wg *sync.WaitGroup, resultChan chan<- responseProcess1) {
	defer wg.Done()

	// Simulasi pemrosesan
	data := 10
	result := data * 2

	// Mengirimkan hasil ke channel
	resultChan <- responseProcess1{Data: result, Error: nil}
}

func processData2(wg *sync.WaitGroup, resultChan chan<- responseProcess2) {
	defer wg.Done()

	// Simulasi pemrosesan
	data := "Hello"
	result := data + " World"

	// Mengirimkan hasil ke channel
	resultChan <- responseProcess2{Data: result, Error: nil}
}

func main() {
	var wg sync.WaitGroup

	// Membuat channel untuk menerima hasil dengan buffer 1
	resultChan1 := make(chan responseProcess1, 1)
	resultChan2 := make(chan responseProcess2, 1)

	wg.Add(2)

	// Running Go Routines
	go processData1(&wg, resultChan1)
	go processData2(&wg, resultChan2)

	wg.Wait()

	result1 := <-resultChan1
	result2 := <-resultChan2

	fmt.Println("Result 1:", result1.Data)
	fmt.Println("Result 2:", result2.Data)
}
