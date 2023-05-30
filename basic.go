package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Memulai Tugas %d\n", taskID)

	time.Sleep(2 * time.Second)

	fmt.Printf("Selesai Tugas %d\n", taskID)
}

func main() {
	var wg sync.WaitGroup

	numTasks := 5

	// Menambahkan jumlah goroutine yang akan dijalankan
	wg.Add(numTasks)

	for i := 1; i <= numTasks; i++ {
		go processTask(i, &wg)
	}

	wg.Wait()

	fmt.Println("All Done!")
}
