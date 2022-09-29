package finn_test

import (
	"encoding/json"
	"fmt"
	"github.com/kaiaverkvist/go-finn-api"
	"github.com/kaiaverkvist/go-finn-api/clients"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

var (
	searchUrls = []string{
		"https://www.finn.no/car/used/search.html?body_type=2&body_type=9&body_type=7&body_type=6&body_type=3&condition=1&exterior_colour=9&exterior_colour=2&fuel=4&fuel=1&fuel=2&location=20061&location=20015&location=22046&make=0.801&make=0.796&make=0.795&mileage_from=20000&model=1.8078.2000501&model=1.8078.2000379&model=1.787.7155&price_from=15400&sort=PUBLISHED_DESC&stored-id=58746383&transmission=1&transmission=2&warranty_insurance=42&warranty_insurance=1&wheel_drive=1&wheel_drive=2&wheel_drive=3&wheel_sets=1&wheel_sets=2&year_from=2000",
		"https://www.finn.no/car/used/search.html?exterior_colour=9&model=1.804.2000301&sort=PUBLISHED_DESC&stored-id=58746383",
	}
	multiPageTestUrl          = "https://www.finn.no/car/used/search.html?model=1.784.2000328&sort=PUBLISHED_DESC"
	secondaryMultiPageTestUrl = "https://www.finn.no/realestate/lettings/search.html?area_from=40&area_to=50&location=1.20061.20528&location=1.20061.20507&location=1.20061.20519&location=1.20061.20512&location=1.20061.20529&location=1.20061.20511&location=1.20061.20523&location=1.20061.20522&location=1.20061.20518&location=1.20061.20513&location=1.20061.20510&location=1.20061.20530&location=1.20061.20509&location=1.20061.20533&location=1.20061.20508&location=1.20061.20531&location=1.20061.20521&page=2&property_type=3&sort=PUBLISHED_DESC&stored-id=58974941"
)

type MockClient struct{}

func (m MockClient) GenericSearch(uri string) (*finn.GenericSearchResult, error) {
	jsonFile, err := os.Open("testfiles/car-ads-response.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var result finn.GenericSearchResult

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m MockClient) RealEstateSearch(uri string) (*finn.SearchResult[finn.RealEstateListing], error) {
	jsonFile, err := os.Open("testfiles/real-estate-response-2.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var result finn.SearchResult[finn.RealEstateListing]

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m MockClient) CarSearch(uri string) (*finn.SearchResult[finn.CarListing], error) {
	jsonFile, err := os.Open("testfiles/car-ads-response.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var result finn.SearchResult[finn.CarListing]

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func TestNewCarSearch_ValidSearchURL(t *testing.T) {
	for _, url := range searchUrls {
		s, err := finn.NewCarSearch(url, MockClient{})

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

func TestNewCarSearch_FetchAds(t *testing.T) {
	for _, url := range searchUrls {
		s, err := finn.NewCarSearch(url, MockClient{})
		assert.Nil(t, err)
		assert.NotNil(t, s)
		s.FetchAds(false)
	}
}

func TestNewCarSearch_RealClient(t *testing.T) {
	url := searchUrls[0]
	s, err := finn.NewCarSearch(url, clients.NewWebsiteClient())
	assert.Nil(t, err)
	assert.NotNil(t, s)
	_, err = s.FetchAds(false)
	assert.Nil(t, err)
}

func TestNewCarSearch_MultiPageTestUrl_RealClient(t *testing.T) {
	url := multiPageTestUrl
	s, err := finn.NewCarSearch(url, clients.NewWebsiteClient())
	if err != nil {
		log.Println(err)
	}
	s.SetDelay(time.Millisecond * 200)
	_, err = s.FetchAds(true)
	assert.Nil(t, err)
}
