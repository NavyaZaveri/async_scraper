package scraper

import (
	"io/ioutil"
	"net/http"
)

func extractBytesFrom(url string)[]byte {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}
	return body

}
