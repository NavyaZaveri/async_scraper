package scraper

import (
	"encoding/json"
	"strconv"
	"testing"
)

func Xkcd(i int) string {
	return "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
}

type XkcdResp struct {
	Img string
}

type XkcdIterator struct {
	curPage  string
	curCount int
}

func (b *XkcdIterator) Next() string {
	b.curPage = Xkcd(b.curCount)
	b.curCount += 1
	return b.curPage
}

func (b *XkcdIterator) HasNext() bool {

	//want just  50
	if b.curCount == 200 {
		return false
	}
	return true
}

func BenchmarkNewWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := NewWorkerPool(40).Fetch(&XkcdIterator{})

		for _, htmlBody_ := range v {
			js := XkcdResp{}
			err := json.Unmarshal(htmlBody_, &js)
			if err == nil {
				//fmt.Println(js)

			} else {
				//fmt.Println(err)
			}

		}
	}

}

func BenchmarkSynchronousScraper(b *testing.B) {

	for i := 0; i < b.N; i++ {
		it := &XkcdIterator{}

		for it.HasNext() {
			page := it.Next()
			body := extractBytesFrom(page)
			js := XkcdResp{}
			err := json.Unmarshal(body, &js)
			if err == nil {
				//fmt.Println(js.Img)
			} else {
				//t.Log(err)
			}
		}
	}
}
