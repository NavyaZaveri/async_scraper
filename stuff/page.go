package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Xkcd(i int) string {
	return "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
}

type Response interface {
}

type xkcdResp struct {
	Img string `json:"img"`
}

type PageIterator interface {
	Next() string
	HasNext() bool
}

type blah struct {
	curPage  string
	curCount int
}

func (b *blah) HasNext() bool {
	if b.curCount == 5 {
		return false
	}
	return true
}

func (b *blah) Next() string {
	b.curPage = Xkcd(b.curCount)
	b.curCount += 1
	return b.curPage
}

func get_resp(url string, k Response) (Response) {
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	r := k
	err = json.Unmarshal(body, &r)
	if err == nil {
		return r
	}
	return nil

}

func run(t PageIterator) {
	c := &xkcdResp{}
	arr := []Response{}
	for t.HasNext() {
		arr = append(arr, get_resp(t.Next(), c))
	}

	for _, v := range arr {
		orig, ok := v.(*xkcdResp)
		fmt.Println(ok)
		if ok {
			fmt.Println("WWU")
			fmt.Println(orig.Img)
		}
	}

}
/*
func main() {
	b := &blah{}
	run(b)
	fmt.Println("DONE/??")

}*/
