package finn

import "time"

type Listing struct {
	Type          string `json:"type"`
	AdId          int    `json:"ad_id"`
	MainSearchKey string `json:"main_search_key"`
	Heading       string `json:"heading"`
	Location      string `json:"location"`
	Image         struct {
		Url    string `json:"url"`
		Path   string `json:"path"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"image"`
	Flags       []string `json:"flags"`
	Timestamp   int64    `json:"timestamp"`
	Coordinates struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coordinates"`
	AdType int    `json:"ad_type"`
	AdLink string `json:"ad_link"`
}

type CarListing struct {
	Listing

	Price struct {
		Amount       int    `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price"`
	OrganisationName string   `json:"organisation_name"`
	Year             int      `json:"year"`
	Mileage          int      `json:"mileage"`
	DealerSegment    string   `json:"dealer_segment"`
	WarrantyDuration int      `json:"warranty_duration"`
	ServiceDocuments []string `json:"service_documents"`
	RegNo            string   `json:"regno"`
}

type RealEstateListing struct {
	Listing

	Styling []string `json:"styling"`
	Logo    struct {
		Url  string `json:"url"`
		Path string `json:"path"`
	} `json:"logo"`
	PriceSuggestion struct {
		Amount       int    `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price_suggestion"`
	PriceTotal struct {
		Amount       int    `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price_total"`
	PriceSharedCost struct {
		Amount       int    `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price_shared_cost"`
	AreaRange struct {
		SizeFrom    int    `json:"size_from"`
		SizeTo      int    `json:"size_to"`
		Unit        string `json:"unit"`
		Description string `json:"description"`
	} `json:"area_range"`
	AreaPlot struct {
		Size        int    `json:"size"`
		Unit        string `json:"unit"`
		Description string `json:"description"`
	} `json:"area_plot"`
	OrganisationName        string      `json:"organisation_name"`
	NumberOfBedrooms        int         `json:"number_of_bedrooms"`
	OwnerTypeDescription    string      `json:"owner_type_description"`
	PropertyTypeDescription string      `json:"property_type_description"`
	ViewingTimes            []time.Time `json:"viewing_times"`

	ImageUrls []string `json:"image_urls"`
}
