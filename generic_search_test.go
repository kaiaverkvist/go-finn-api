package finn_test

import (
	"github.com/kaiaverkvist/go-finn-api"
	"github.com/kaiaverkvist/go-finn-api/clients"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestNewGenericSearch_ValidSearchURL(t *testing.T) {
	for _, url := range searchUrls {
		s, err := finn.NewGenericSearch(url, MockClient{})

		// Can we create a search object given these urls...
		assert.Nil(t, err, "expected successful search object creation")
		assert.NotNil(t, s, "expected search object to be created")

		// Do we have a URL object?
		assert.NotNil(t, s.ParsedUri, "expected non nil url object")

		// Do we have attributes?
		assert.NotEmpty(t, s.Attributes, "expected attributes to not be empty")

		assert.Equalf(t, url, s.OriginalUri, "original url does not match")
	}
}

func TestSearch_FetchAds(t *testing.T) {
	for _, url := range searchUrls {
		s, err := finn.NewGenericSearch(url, MockClient{})
		assert.Nil(t, err)
		assert.NotNil(t, s)
		s.FetchAds(false)
	}
}

func TestNewGenericSearch_RealClient(t *testing.T) {
	url := searchUrls[0]
	s, err := finn.NewGenericSearch(url, clients.NewWebsiteClient())
	assert.Nil(t, err)
	assert.NotNil(t, s)
	_, err = s.FetchAds(false)
	assert.Nil(t, err)
}

func TestNewGenericSearch_MultiPageTestUrl_RealClient(t *testing.T) {
	url := secondaryMultiPageTestUrl
	s, err := finn.NewGenericSearch(url, clients.NewWebsiteClient())
	if err != nil {
		log.Println(err)
	}
	s.SetDelay(time.Millisecond * 200)
	ads, err := s.FetchAds(true)
	assert.NotNil(t, ads)
	assert.Nil(t, err)
}
