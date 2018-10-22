package testdata

import "strconv"

//Xkcd returns the url to scrape
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

	//want just  200 comics
	if b.curCount == 200 {
		return false
	}
	return true
}