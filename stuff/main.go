package main

import (
	"fmt"
	"sort"
	"sync"
)

/*
implementing a worker pool system
*/

func worker(jobs <-chan int, result chan<- int) {
	for i := range jobs {
		result <- i + 10
	}
}

func main() {
	jobs := make(chan int)
	result := make(chan int)
	for i := 0; i < 2; i++ {
		go worker(jobs, result)
	}
	wg := sync.WaitGroup{}
	res := []int{}


	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(data int) {
			defer wg.Done()
			jobs <- data
			x := <-result
			res = append(res, x)
		}(i)
	}

	wg.Wait()
	close(jobs)
	sort.Slice(res, func(i, j int) bool {
		return res[i]<res[j]
	})
	fmt.Println(res)

}
