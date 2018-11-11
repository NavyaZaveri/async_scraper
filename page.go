package scraper


/*
PageIterator is an iterable interface with two methods.

	Next() should return the url to be scraped

	HasNext() should return a bool if there are any
	more pages to scrape

*/
type PageIterator interface {


	Next() string
	HasNext() bool
}
