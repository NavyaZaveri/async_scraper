package scraper

import "testing"


//TODO
func fetch(it *XkcdIterator) [][]byte {
	res := [][]byte{}

	for it.HasNext() {

	}
	return res
}

func TestSync(t *testing.T) {
	_ = NewWorkerPool(100).Fetch(&XkcdIterator{})
}

