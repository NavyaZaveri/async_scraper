package scraper

type PageIterator interface {

	/*
		Next() should return the url to be scraped

		HasNext() should return a bool if there are any
		more pages to scrape

	*/

	Next() string
	HasNext() bool
}
