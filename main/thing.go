package main

import (
	"fmt"
	"github.com/NavyaZaveri/scraper"
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

func main() {

	v := scraper.NewWorkerPool(20).Fetch(&scraper.XkcdIterator{})
	fmt.Println(len(v))

}
