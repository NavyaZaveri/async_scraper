package main

import (
	"fmt"
	"sync"
)

/*
worker pool implementation
so you have workers and jobs
each worker is assigned  some job, concurrently


ideal api
w worker(2).fetch()

new WorkedPool().setJob().fetch()

*/

type Worker struct {
}

type Jobs chan Job

type Job string

type WorkerPool struct {
	workers []*Worker
	result  chan string
}

func (w *WorkerPool) fetch(p PageIterator) []string {
	jobs := make(Jobs, 0)
	w.spawnWorkers(jobs)
	wg := sync.WaitGroup{}
	results := []string{}
	for p.HasNext() {
		currentPage := p.Next()
		wg.Add(1)
		go func(cur string) {
			defer wg.Done()
			jobs <- Job(cur)

			x := <-w.result
			fmt.Println(x)
			results = append(results, x)
		}(currentPage)
	}
	wg.Wait()
	close(jobs)
	return results
}

func (w *WorkerPool) spawnWorkers(j Jobs) {
	for _, worker := range w.workers {
		go work(worker, j, w.result)
	}
}
func (w *Worker) execute(j Job) string {
	return string(j)

}

func newWorker() *Worker {
	return &Worker{}
}

func NewWorkerPool(num int) *WorkerPool {
	w := &WorkerPool{result: make(chan string, 0)}
	for i := 0; i < num; i++ {
		w.workers = append(w.workers, newWorker())
	}

	return w
}

func work(w *Worker, jobs Jobs, res chan<- string) {
	for job := range jobs {
		ans := w.execute(job)

		//blocks
		res <- ans
	}
}

func main() {
	b := &blah{}
	//NewWorkerPool(3).Fetch("HELLO", 10)
	NewWorkerPool(3).fetch(b)
	fmt.Println("done")
}
