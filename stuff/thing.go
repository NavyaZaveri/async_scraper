package main

import "fmt"

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
	result  chan int
}

func (w *WorkerPool) Fetch(job Job, howMany int) {
	jobs := make(Jobs, 0)
	w.spawnWorkers(jobs)

	for i := 0; i < howMany; i++ {
		jobs <- "WEO"
		x := <-w.result
		fmt.Println(x)
	}
	close(jobs)
}

func (w *WorkerPool) spawnWorkers(j Jobs) {
	for _, worker := range w.workers {
		go work(worker, j, w.result)
	}
}
func (w *Worker) execute(j Job) int {
	return 1

}

func newWorker() *Worker {
	return &Worker{}
}

func NewWorkerPool(num int) *WorkerPool {
	w := &WorkerPool{result: make(chan int, 0)}
	for i := 0; i < num; i++ {
		w.workers = append(w.workers, newWorker())
	}

	return w
}

func work(w *Worker, jobs Jobs, res chan<- int) {
	for job := range jobs {
		ans := w.execute(job)

		//blocks
		res <- ans
	}
}

func main() {

	NewWorkerPool(3).Fetch("HELLO", 10)
	fmt.Println("done")

}
