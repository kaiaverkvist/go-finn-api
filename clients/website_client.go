package clients

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/kaiaverkvist/go-finn-api"
)

type WebsiteClient struct {
	c *colly.Collector
}

func (w WebsiteClient) RealEstateSearch(uri string) (*finn.SearchResult[finn.RealEstateListing], error) {
	var result finn.SearchResult[finn.RealEstateListing]
	var err error

	// Extract inner html from the __NEXT_DATA__ tagged
	w.c.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		x := e.Text
		err = json.Unmarshal([]byte(x), &result)
	})

	if err != nil {
		return nil, err
	}

	err = w.c.Visit(uri)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (w WebsiteClient) CarSearch(uri string) (*finn.SearchResult[finn.CarListing], error) {
	var result finn.SearchResult[finn.CarListing]
	var err error

	// Extract inner html from the __NEXT_DATA__ tagged
	w.c.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		x := e.Text
		err = json.Unmarshal([]byte(x), &result)
	})

	if err != nil {
		return nil, err
	}

	err = w.c.Visit(uri)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (w WebsiteClient) GenericSearch(uri string) (*finn.GenericSearchResult, error) {
	var result finn.GenericSearchResult
	var err error

	// Extract inner html from the __NEXT_DATA__ tagged
	w.c.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		x := e.Text
		err = json.Unmarshal([]byte(x), &result)
	})

	if err != nil {
		return nil, err
	}

	err = w.c.Visit(uri)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewWebsiteClient() *WebsiteClient {
	return &WebsiteClient{c: colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)}
}
