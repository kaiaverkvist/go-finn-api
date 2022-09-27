package finn

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	searchType_Generic = "generic"
)

func NewGenericSearch(uri string, client IClient) (*GenericSearch, error) {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to url parsing error: %s", err.Error())
	}

	params, err := url.ParseQuery(parsedUri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("could not create search type due to query params parsing error: %s", err.Error())
	}

	search := GenericSearch{
		searchType:   searchType_Generic,
		client:       client,
		OriginalUri:  uri,
		ParsedUri:    parsedUri,
		Attributes:   params,
		requestDelay: time.Second * 2,
	}

	// Attributes parsing (query params to map)
	return &search, nil
}

type GenericSearch struct {
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

func (s *GenericSearch) SetDelay(duration time.Duration) {
	s.requestDelay = duration
}

func (s *GenericSearch) FetchAds(iteratePages bool) ([]map[string]any, error) {

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
	var docs []map[string]any
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

func (s *GenericSearch) GetTitle() string {
	return s.searchTitle
}

func (s *GenericSearch) iteratePage(page int) ([]map[string]any, error) {
	var docs []map[string]any
	// Initial search result for the first page in the search.
	uri := *s.ParsedUri
	uri.Query().Set("page", fmt.Sprintf("%d", page))

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
