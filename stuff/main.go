package main

import "fmt"

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
	for i := 0; i < 10; i++ {
		go worker(jobs, result)
	}

	for  i:=0;i<20;i++ {
		jobs <- i

		x := <-result
		fmt.Println(x)

	}

	close(jobs)

}
