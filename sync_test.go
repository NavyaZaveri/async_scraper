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

	//use a map to store unique keys
	set := map[testdata.XkcdResp]int{}

	for _, xkcdBytes := range a1 {
		err := json.Unmarshal(xkcdBytes, &x)
		if err == nil {
			set[x] = 1
			comics = append(comics, x)
		}
	}
	if len(set) != len(comics) {
		t.Logf("TestForDuplicates Failed")
	} else {
		t.Logf("test passed")
	}

}
