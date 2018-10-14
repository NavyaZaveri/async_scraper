package testdata

import "strconv"

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