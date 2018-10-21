package scraper

import (
	"sync"
)

type WorkerPool struct {
	workers []*Worker
	result  chan []byte
}

//TODO:  set hard limit or throw error if iteration count not updates

//NewWorkerPool returns an array of workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	w := &WorkerPool{result: make(chan []byte, 0)}

	for i := 0; i < numWorkers; i++ {
		w.workers = append(w.workers, newWorker())
	}
	return w
}

func (w *WorkerPool) spinWorkers(j JobQueue) {
	for _, worker := range w.workers {
		go worker.work(j, w.result)
	}
}

/*
Fetch returns a 2-d slice, where each element is the byte contents of each website
needed to scrape
*/
func (w *WorkerPool) Fetch(p PageIterator) [][]byte {

	jobQueue := make(JobQueue, 0)

	//each worker is now "waiting" for a job
	w.spinWorkers(jobQueue)

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	results := [][]byte{}

	for p.HasNext() {
		currentPage := p.Next()
		wg.Add(1)

		go func(cur string) {
			defer wg.Done()

			//add a job to the channel queue. A worker
			//will pick this job up when it's free. If there are multiple
			//free workers, a worker is selected pseudorandomly
			jobQueue <- Job(cur)

			//get the result of the job.
			x := <-w.result

			//lock to prevent race conditions
			mux.Lock()

			//store the (non-null) result
			if x != nil {
				results = append(results, x)
			}

			mux.Unlock()

		}(currentPage)
	}
	wg.Wait()

	//close the channel so workers stop waiting
	close(jobQueue)
	return results
}
