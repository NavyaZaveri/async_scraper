package scraper

import (
	"sync"
)

type WorkerPool struct {
	workers []*Worker
	result  chan []byte
}

const PAGE_ITERATOR_LIMT int = 1000

func NewWorkerPool(numWorkers int) *WorkerPool {
	w := &WorkerPool{result: make(chan []byte, 0)}

	for i := 0; i < numWorkers; i++ {
		w.workers = append(w.workers, newWorker())
	}
	return w
}

func (w *WorkerPool) spinWorkers(j Jobs) {
	for _, worker := range w.workers {
		go work(worker, j, w.result)
	}
}

func (w *WorkerPool) Fetch(p PageIterator) [][]byte {
	jobs := make(Jobs, 0)

	//each worker is now "waiting" for a job
	w.spinWorkers(jobs)

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	results := [][]byte{}

	for p.HasNext() {
		currentPage := p.Next()
		wg.Add(1)

		go func(cur string) {
			defer wg.Done()

			//add a job to the channel queue
			jobs <- Job(cur)

			//get the result of some job.
			x := <-w.result

			//store the result
			mux.Lock()

			if x != nil {
				results = append(results, x)
			}

			mux.Unlock()

		}(currentPage)
	}
	wg.Wait()

	//close the channel so workers stop waiting
	close(jobs)
	return results
}
