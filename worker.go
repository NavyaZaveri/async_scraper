package scraper

import (
	"io/ioutil"
	"net/http"
)

type Worker struct {
}

func newWorker() *Worker {
	return &Worker{}
}

func (w *Worker) execute(j Job) []byte {
	url := string(j)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return body
	}
	return nil

}

func work(w *Worker, jobs Jobs, res chan<- []byte) {
	for job := range jobs {

		//you have one job
		ans := w.execute(job)

		//send the result to the channel
		res <- ans
	}
}
