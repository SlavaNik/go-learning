package main

import (
	"fmt"
	"go_lect/internal/worker_pool"
)

func main() {

	wc := func() worker_pool.Worker {
		return TestWorker{}
	}
	wpool := worker_pool.NewPool(4, wc)
	wpool.DataSource = make(chan interface{})

	go wpool.Run()
	testStrSlice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, t := range testStrSlice {
		wpool.DataSource <- t
	}

	wpool.Stop()
}

type TestWorker struct{}

func (w TestWorker) Do(data interface{}, i int) {
	fmt.Printf("Goroutine number %d, is running. Input %s \n", i, data)
}
func (w TestWorker) Stop() {

}
