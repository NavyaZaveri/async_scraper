package scraper

import (
	"encoding/json"
	"github.com/NavyaZaveri/scraper/testdata"
	"testing"
)

func BenchmarkNewWorkerPool(b *testing.B) {

	for i := 0; i < b.N; i++ {
		v := NewWorkerPool(100).Fetch(&testdata.XkcdIterator{})

		for _, htmlBody_ := range v {
			xkcdJson := testdata.XkcdResp{}
			_ = json.Unmarshal(htmlBody_, &xkcdJson)

		}
	}
}

func BenchmarkSynchronousScraper(b *testing.B) {

	for i := 0; i < b.N; i++ {
		it := &testdata.XkcdIterator{}

		for it.HasNext() {
			page := it.Next()
			body := extractBytesFrom(page)
			xkcdJson := testdata.XkcdResp{}
			_ = json.Unmarshal(body, &xkcdJson)
		}
	}
}
