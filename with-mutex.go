package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Value int
}

type DataMerger struct {
	result []Data
	mutex  sync.Mutex
}

func (dm *DataMerger) Merge(data Data) {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	dm.result = append(dm.result, data)
}

func main() {
	var wg sync.WaitGroup

	merger := DataMerger{}

	numProcess := 5

	wg.Add(numProcess)

	for i := 0; i < numProcess; i++ {
		go func(id int) {
			defer wg.Done()

			data := Data{
				Value: id + 1,
			}

			merger.Merge(data)
		}(i)
	}

	wg.Wait()

	fmt.Println("Result :")

	fmt.Println(merger.result)
	for _, data := range merger.result {
		fmt.Println(data.Value)
	}
}
