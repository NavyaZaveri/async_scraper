package scraper

type Worker struct {
}

func newWorker() *Worker {
	return &Worker{}
}

func (w *Worker) execute(j Job) []byte {
	return extractBytesFrom(string(j))
}

func work(w *Worker, jobs Jobs, res chan<- []byte) {
	for job := range jobs {

		//you have one job. Do it!
		ans := w.execute(job)

		//sends the result to the channel.
		// The result is received from the other
		//end, freeing up the worker.
		res <- ans
	}
}
