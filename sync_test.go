package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/NavyaZaveri/scraper/testdata"
	"testing"
)

//TODO

func fetch(it *testdata.XkcdIterator) [][]byte {
	res := [][]byte{}

	for it.HasNext() {
		jsonBytes := extractBytesFrom(it.Next())
		res = append(res, jsonBytes)

	}
	return res
}

func TestSync(t *testing.T) {
	a1 := NewWorkerPool(100).Fetch(&testdata.XkcdIterator{})
	x := testdata.XkcdResp{}
	mmap := map[string]int{}
	for _, xkcdJson := range a1 {
		err := json.Unmarshal(xkcdJson, &x)
		mmap[x.Img] = 1
		if err!=nil {
			fmt.Println("WARNIGN"+err.Error())
		}
	}

	if len(mmap) != len(a1) {
		t.Logf("test failed")
	} else{
		t.Logf("test passed")
	}

}
