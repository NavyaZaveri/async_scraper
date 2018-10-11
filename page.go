package scraper

import (
	"strconv"
)

func Xkcd(i int) string {
	return "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
}

type XkcdResp struct {
	Img string `json:"img"`
}

type PageIterator interface {
	Next() string
	HasNext() bool
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
