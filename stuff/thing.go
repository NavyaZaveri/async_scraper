package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/*
worker pool implementation
so you have workers and jobs
each worker is assigned  some job, concurrently

CHANGE OF PLANS
do **not** unmarshal, just send the body
let the user unmarshal

ideal api

NeWorkePool().setJob().fetch()

*/

type Worker struct {
}

type Jobs chan Job

type Job string

type WorkerPool struct {
	workers []*Worker
	result  chan []byte
}

func (w *WorkerPool) fetch(p PageIterator) [][]byte {
	jobs := make(Jobs, 0)
	w.spawnWorkers(jobs)
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	results := [][]byte{}
	for p.HasNext() {
		currentPage := p.Next()
		wg.Add(1)

		//assign a job to  a
		//worker asynchronously
		go func(cur string) {
			defer wg.Done()
			fmt.Println("CUR IS" + cur)
			jobs <- Job(cur)

			//unblocks

			x := <-w.result
			mux.Lock()
			results = append(results, x)

			mux.Unlock()
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

func (w *Worker) execute(j Job) []byte {
	url := string(j)
	fmt.Println(url)
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)


	if err == nil {
		return body
	}
	return nil

}

func newWorker() *Worker {
	return &Worker{}
}

func NewWorkerPool(num int) *WorkerPool {
	w := &WorkerPool{result: make(chan []byte, 0)}
	for i := 0; i < num; i++ {
		w.workers = append(w.workers, newWorker())
	}
	return w
}

func work(w *Worker, jobs Jobs, res chan<- []byte) {
	for job := range jobs {
		fmt.Println("job is" + job)
		ans := w.execute(job)

		res <- ans
	}
}

func main() {

	v := NewWorkerPool(20).fetch(&blah{})
	fmt.Println(v)
	x := xkcdResp{}
	sample := v[2]
	err := json.Unmarshal(sample, &x)
	if err==nil {
		fmt.Println(x.Img)

	} else {
		panic(err)
	}


	}



