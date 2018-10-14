package main

import (
	"encoding/json"
	"fmt"
	"github.com/NavyaZaveri/scraper"
	"strconv"
)
/*
ideal usage?
NeWorkerPool().fetch()

*/

func xkcd() {

}

type XkcdResp struct {
	Img string
}


type XkcdIterator struct {
	curPage  string
	curCount int
}

func (b *XkcdIterator) HasNext() bool {

	//want just  50
	if b.curCount == 50 {
		return false
	}
	return true
}



func (b *XkcdIterator) Next() string {
	b.curPage = Xkcd(b.curCount)
	b.curCount += 1
	return b.curPage
}


func main() {

	//spin up twenty workers to fetch stuff from the links
	//provided by the iterator
	v := scraper.NewWorkerPool(40).Fetch(&XkcdIterator{})

	for _, htmlBody_ := range v {
		js := XkcdResp{}
		err := json.Unmarshal(htmlBody_, &js)
		if err == nil {
			fmt.Println(js)

		} else {
			fmt.Println(err)
		}

	}

}
