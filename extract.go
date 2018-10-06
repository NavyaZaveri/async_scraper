package scraper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func extract(xkcdUrl string) (map[string]interface{}, error) {
	resp, err := http.Get(xkcdUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	m := map[string]interface{}{}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(contents), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

