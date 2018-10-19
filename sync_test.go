package scraper

import (
	"encoding/json"
	"github.com/NavyaZaveri/scraper/testdata"
	"testing"
)


func TestForDuplicates(t *testing.T) {
	a1 := NewWorkerPool(100).Fetch(&testdata.XkcdIterator{})
	comics := []testdata.XkcdResp{}
	x := testdata.XkcdResp{}
	mmap := map[testdata.XkcdResp]int{}
	for _, xkcdJson := range a1 {
		err := json.Unmarshal(xkcdJson, &x)
		if err == nil {
			mmap[x] = 1
			comics = append(comics,x)
		} else {
		}
	}

	if len(mmap) != len(comics) {
		t.Logf("TestForDuplicates Failed")
	} else{
		t.Logf("test passed")
	}

}
