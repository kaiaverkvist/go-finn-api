package finn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	searchType_Car        = "car"
	searchType_RealEstate = "realestate"
)

type Search[T any] struct {
	client IClient

	searchType string

	OriginalUri string
	ParsedUri   *url.URL

	Attributes url.Values

	totalResults  int
	pageCount     int
	iteratedPages int
	searchTitle   string

	requestDelay time.Duration
}

func NewCarSearch(uri string, client IClient) (*Search[CarListing], error) {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to url parsing error: %s", err.Error())
	}

	params, err := url.ParseQuery(parsedUri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to query params parsing error: %s", err.Error())
	}

	search := Search[CarListing]{
		searchType:   searchType_Car,
		client:       client,
		OriginalUri:  uri,
		ParsedUri:    parsedUri,
		Attributes:   params,
		requestDelay: time.Second * 2,
	}

	// Attributes parsing (query params to map)
	return &search, nil
}

func NewRealEstateSearch(uri string, client IClient) (*Search[RealEstateListing], error) {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to url parsing error: %s", err.Error())
	}

	params, err := url.ParseQuery(parsedUri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to query params parsing error: %s", err.Error())
	}

	search := Search[RealEstateListing]{
		searchType:   searchType_RealEstate,
		client:       client,
		OriginalUri:  uri,
		ParsedUri:    parsedUri,
		Attributes:   params,
		requestDelay: time.Second * 2,
	}

	// Attributes parsing (query params to map)
	return &search, nil
}

func (s *Search[T]) SetDelay(duration time.Duration) {
	s.requestDelay = duration
}

func (s *Search[T]) FetchAds(iteratePages bool) ([]T, error) {

	// TODO: Get rid of this switch statement
	switch s.searchType {
	case searchType_Generic:
		{
			// Initial search result for the first page in the search.
			firstResult, err := s.client.GenericSearch(s.OriginalUri)
			if err != nil {
				return nil, err
			}

			s.pageCount = firstResult.Props.PageProps.Search.Metadata.Paging.Last
			s.searchTitle = firstResult.Props.PageProps.Search.ResultHeading
			s.totalResults = firstResult.Props.PageProps.Search.Metadata.ResultSize.MatchCount

			time.Sleep(s.requestDelay)
			s.iteratedPages = 1
			var docs []T
			if iteratePages {
				for i := 0; i < s.pageCount; i++ {
					d, err := s.iteratePage(i + 1)
					if err != nil {
						return nil, err
					}

					docs = append(docs, d...)
					s.iteratedPages += 1
					time.Sleep(s.requestDelay)
				}
			}

			return docs, nil
		}
	case searchType_Car:
		{
			// Initial search result for the first page in the search.
			firstResult, err := s.client.CarSearch(s.OriginalUri)
			if err != nil {
				return nil, err
			}

			s.pageCount = firstResult.Props.PageProps.Search.Metadata.Paging.Last
			s.searchTitle = firstResult.Props.PageProps.Search.ResultHeading
			s.totalResults = firstResult.Props.PageProps.Search.Metadata.ResultSize.MatchCount

			time.Sleep(s.requestDelay)
			s.iteratedPages = 1
			var docs []T
			if iteratePages {
				for i := 0; i < s.pageCount; i++ {
					d, err := s.iteratePage(i + 1)
					if err != nil {
						return nil, err
					}

					docs = append(docs, d...)
					s.iteratedPages += 1
					time.Sleep(s.requestDelay)
				}
			}

			return docs, nil
		}
	case searchType_RealEstate:
		{
			// Initial search result for the first page in the search.
			firstResult, err := s.client.RealEstateSearch(s.OriginalUri)
			if err != nil {
				return nil, err
			}

			s.pageCount = firstResult.Props.PageProps.Search.Metadata.Paging.Last
			s.searchTitle = firstResult.Props.PageProps.Search.ResultHeading
			s.totalResults = firstResult.Props.PageProps.Search.Metadata.ResultSize.MatchCount

			time.Sleep(s.requestDelay)
			s.iteratedPages = 1
			var docs []T
			if iteratePages {
				for i := 0; i < s.pageCount; i++ {
					d, err := s.iteratePage(i + 1)
					if err != nil {
						return nil, err
					}

					docs = append(docs, d...)
					s.iteratedPages += 1
					time.Sleep(s.requestDelay)
				}
			}

			return docs, nil
		}
	}

	return nil, errors.New("invalid search type used")
}

func (s *Search[T]) GetTitle() string {
	return s.searchTitle
}

func (s *Search[T]) iteratePage(page int) ([]T, error) {
	var docs []T
	// Initial search result for the first page in the search.
	uri := *s.ParsedUri
	uri.Query().Set("page", fmt.Sprintf("%d", page))

	// TODO: Get rid of this switch statement
	switch s.searchType {
	case searchType_Generic:
		{
			result, err := s.client.GenericSearch(uri.String())
			if err != nil {
				return nil, err
			}

			resultDocs, err := json.Marshal(result.Props.PageProps.Search.Docs)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(resultDocs, &docs)
			if err != nil {
				return nil, err
			}

			return docs, nil
		}
	case searchType_Car:
		{
			result, err := s.client.CarSearch(uri.String())
			if err != nil {
				return nil, err
			}

			resultDocs, err := json.Marshal(result.Props.PageProps.Search.Docs)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(resultDocs, &docs)
			if err != nil {
				return nil, err
			}

			return docs, nil
		}
	case searchType_RealEstate:
		{
			result, err := s.client.RealEstateSearch(uri.String())
			if err != nil {
				return nil, err
			}

			resultDocs, err := json.Marshal(result.Props.PageProps.Search.Docs)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(resultDocs, &docs)
			if err != nil {
				return nil, err
			}

			return docs, nil
		}
	}

	return nil, errors.New("invalid search type used")
}
