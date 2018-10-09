package stuff

import (
	"fmt"
	"github.com/NavyaZaveri/scraper/xkcd"
	"strconv"
)

func Xkcd(i int) string {
	return "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
}

//fetches howMany comics from the xkcd api
func Fetch(howMany int) xkcd.Comics {
	comics := xkcd.Comics{}
	for i := 0; i < howMany; i++ {
		xkcdUrl := Xkcd(i)
		mmap, err := extract(xkcdUrl)
		if err != nil {
			fmt.Printf("[Warning] Could not fetch comic %d,  %s\n", i, err)
		} else {
			text := xkcd.Parse(mmap["transcript"].(string))
			url := mmap["img"].(string)
			comics = append(comics, xkcd.Comic{Text: text, ImgUrl: url})
		}
	}
	return comics
}

func AsyncFetch(howMany int) xkcd.Comics {
	chanList := make([]chan xkcd.Comic, 0)
	coms := xkcd.Comics{}

	for i := 0; i < howMany; i++ {
		chanList = append(chanList, make(chan xkcd.Comic))

		go func(ch chan<- xkcd.Comic, urlIndex int) {

			//get the xkcd url
			xkcdUrl := Xkcd(urlIndex)

			//extract out the json in the form of a map
			mmap, err := extract(xkcdUrl)
			if err != nil {
				fmt.Printf("[Warning] Could not fetch comic %d,  %s\n", urlIndex, err)
			} else {

				//get the comic text
				text := xkcd.Parse(mmap["transcript"].(string))
				imgUrl := mmap["img"].(string)

				//send a comic object to channel
				ch <- xkcd.Comic{ImgUrl: imgUrl, Text: text}
			}

			close(ch)

		}(chanList[i], i)
	}
	for _, channel := range chanList {
		for comic := range channel {
			coms = append(coms, comic)
		}
	}
	return coms
}
