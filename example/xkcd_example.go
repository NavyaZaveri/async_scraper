package main

import (
	"encoding/json"
	"fmt"
	"github.com/NavyaZaveri/scraper"
)

/*
ideal usage?
NeWorkerPool().fetch()

*/

type JsonResp struct {
	Img string
}

func main() {

	//spin up twenty workers to fetch stuff from the links
	//provided by the iterator
	v := scraper.NewWorkerPool(40).Fetch(&scraper.XkcdIterator{})

	for _, htmlBody_ := range v {
		js := JsonResp{}
		err := json.Unmarshal(htmlBody_, &js)
		if err == nil {
			fmt.Println(js)

		} else {
			fmt.Println(err)
		}

	}

}
