package main

import (
	"encoding/json"
	"fmt"
	"github.com/NavyaZaveri/scraper"
	"github.com/NavyaZaveri/scraper/testdata"
	"runtime"
)

func main() {

	//spin up 40 workers to fetch contents ([]bytes) from the links
	//provided by the iterator
	v := scraper.NewWorkerPool(10).Fetch(&testdata.XkcdIterator{})

	for _, htmlBody_ := range v {
		js := testdata.XkcdResp{}
		err := json.Unmarshal(htmlBody_, &js)
		if err == nil {
			fmt.Println(js.Img)

		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(runtime.NumGoroutine());
}
