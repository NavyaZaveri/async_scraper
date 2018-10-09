package scraper

import (
	"fmt"
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
	fmt.Println(url)
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		return body
	}
	return nil

}


func work(w *Worker, jobs Jobs, res chan<- []byte) {
	for job := range jobs {
		ans := w.execute(job)

		res <- ans
	}
}


