package finn

type IClient interface {
	CarSearch(uri string) (*SearchResult[CarListing], error)
	RealEstateSearch(uri string) (*SearchResult[RealEstateListing], error)
}

type SearchParam struct {
	key   string
	value string
}

type SearchResult[T any] struct {
	Props struct {
		PageProps struct {
			Search struct {
				Docs    []T `json:"docs"`
				Filters []struct {
					DisplayName string `json:"display_name"`
					Name        string `json:"name"`
					FilterItems []struct {
						DisplayName string `json:"display_name"`
						Name        string `json:"name"`
						Value       string `json:"value"`
						Hits        int    `json:"hits"`
						FilterItems []struct {
							DisplayName string        `json:"display_name"`
							Name        string        `json:"name"`
							Value       string        `json:"value"`
							Hits        int           `json:"hits"`
							FilterItems []interface{} `json:"filter_items"`
							Selected    bool          `json:"selected"`
						} `json:"filter_items"`
						Selected bool `json:"selected"`
					} `json:"filter_items"`
					Type             string `json:"type"`
					ItemDisplayCount int    `json:"item_display_count,omitempty"`
					MinValue         int    `json:"min_value,omitempty"`
					MaxValue         int    `json:"max_value,omitempty"`
					Step             int    `json:"step,omitempty"`
					Unit             string `json:"unit,omitempty"`
					NameFrom         string `json:"name_from,omitempty"`
					NameTo           string `json:"name_to,omitempty"`
					ValueFrom        int    `json:"value_from,omitempty"`
					IsYear           bool   `json:"is_year,omitempty"`
				} `json:"filters"`
				Metadata struct {
					Params          map[string]any
					SearchKey       string `json:"search_key"`
					SelectedFilters []struct {
						Parameters []struct {
							ParameterName  string `json:"parameter_name"`
							ParameterValue string `json:"parameter_value"`
						} `json:"parameters"`
						FilterName  string `json:"filter_name"`
						DisplayName string `json:"display_name"`
						Prefix      string `json:"prefix"`
					} `json:"selected_filters"`
					NumResults      int `json:"num_results"`
					QuestTime       int `json:"quest_time"`
					SolrTime        int `json:"solr_time"`
					SolrElapsedTime int `json:"solr_elapsed_time"`
					ResultSize      struct {
						MatchCount int `json:"match_count"`
						GroupCount int `json:"group_count"`
					} `json:"result_size"`
					Paging struct {
						Param   string `json:"param"`
						Current int    `json:"current"`
						Last    int    `json:"last"`
					} `json:"paging"`
					Title                string `json:"title"`
					IsSavableSearch      bool   `json:"is_savable_search"`
					SearchKeyDescription string `json:"search_key_description"`
					Vertical             string `json:"vertical"`
					VerticalDescription  string `json:"vertical_description"`
					Sort                 string `json:"sort"`
					Uuid                 string `json:"uuid"`
					Tracking             struct {
						Object struct {
							SelectionFilters []struct {
								Name    string `json:"name"`
								Value   string `json:"value"`
								ValueId string `json:"valueId,omitempty"`
							} `json:"selectionFilters"`
							SortingType string `json:"sortingType"`
							NumItems    int    `json:"numItems"`
							PageNumber  int    `json:"pageNumber"`
							Layout      string `json:"layout"`
							Type        string `json:"type"`
						} `json:"object"`
						Vertical struct {
							Name        string `json:"name"`
							SubVertical string `json:"subVertical"`
						} `json:"vertical"`
					} `json:"tracking"`
					GuidedSearch struct {
						Suggestions []interface{} `json:"suggestions"`
						Tracking    struct {
							Search struct {
								Items []interface{} `json:"items"`
								Type  string        `json:"@type"`
								Id    string        `json:"@id"`
							} `json:"search"`
							Vertical struct {
								Name        string `json:"name"`
								SubVertical string `json:"subVertical"`
							} `json:"vertical"`
							Name   string `json:"name"`
							Intent string `json:"intent"`
							Type   string `json:"type"`
						} `json:"tracking"`
					} `json:"guided_search"`
					Actions       []interface{} `json:"actions"`
					IsEndOfPaging bool          `json:"is_end_of_paging"`
				} `json:"metadata"`
				MapUrl       string `json:"mapUrl"`
				PageMetadata struct {
					Title          string `json:"title"`
					Description    string `json:"description"`
					IndexDirective string `json:"indexDirective"`
					CanonicalUrl   string `json:"canonicalUrl"`
					OpenGraphUrl   string `json:"openGraphUrl"`
					JsonLd         struct {
						Context           string `json:"@context"`
						Type              string `json:"@type"`
						Url               string `json:"url"`
						MainContentOfPage struct {
							Type        string `json:"@type"`
							CssSelector string `json:"cssSelector"`
						} `json:"mainContentOfPage"`
						Breadcrumb struct {
							Type            string `json:"@type"`
							Name            string `json:"name"`
							ItemListElement []struct {
								Type     string `json:"@type"`
								Position int    `json:"position"`
								Item     struct {
									Id   string `json:"@id"`
									Name string `json:"name"`
								} `json:"item"`
							} `json:"itemListElement"`
						} `json:"breadcrumb"`
						Headline string `json:"headline"`
					} `json:"jsonLd"`
					Image string `json:"image"`
				} `json:"pageMetadata"`
				ResultHeading string `json:"resultHeading"`
			} `json:"search"`
			Banner                 string `json:"banner"`
			RecommendationsContent string `json:"recommendationsContent"`
			ShowRecommendations    bool   `json:"showRecommendations"`
			Layout                 string `json:"layout"`
			SortOptions            []struct {
				Sort        string `json:"sort"`
				Description string `json:"description"`
			} `json:"sortOptions"`
			JobProfileEntryContent    string `json:"jobProfileEntryContent"`
			DisableGridView           bool   `json:"disableGridView"`
			JobRecommendationsContent string `json:"jobRecommendationsContent"`
		} `json:"pageProps"`
		LoginId        string `json:"loginId"`
		DeviceType     string `json:"deviceType"`
		FeatureToggles struct {
			MultiVariantJobRerank struct {
				Enabled bool   `json:"enabled"`
				Variant string `json:"variant"`
			} `json:"multiVariantJobRerank"`
			MultiVariantJobRecommendations struct {
				Enabled bool   `json:"enabled"`
				Variant string `json:"variant"`
			} `json:"multiVariantJobRecommendations"`
			RealestateLettingsMedium bool `json:"realestateLettingsMedium"`
			ShowHeading              bool `json:"showHeading"`
			TjtHotjarEnabled         bool `json:"tjtHotjarEnabled"`
		} `json:"featureToggles"`
		NSSP bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		BodyType          []string `json:"body_type"`
		Condition         string   `json:"condition"`
		ExteriorColour    []string `json:"exterior_colour"`
		Fuel              []string `json:"fuel"`
		Location          []string `json:"location"`
		Make              []string `json:"make"`
		MileageFrom       string   `json:"mileage_from"`
		Model             []string `json:"model"`
		PriceFrom         string   `json:"price_from"`
		Sort              string   `json:"sort"`
		StoredId          string   `json:"stored-id"`
		Transmission      []string `json:"transmission"`
		WarrantyInsurance []string `json:"warranty_insurance"`
		WheelDrive        []string `json:"wheel_drive"`
		WheelSets         []string `json:"wheel_sets"`
		YearFrom          string   `json:"year_from"`
		Vertical          string   `json:"vertical"`
		Subvertical       string   `json:"subvertical"`
	} `json:"query"`
	BuildId       string `json:"buildId"`
	AssetPrefix   string `json:"assetPrefix"`
	RuntimeConfig struct {
		CanonicalBaseUrl string `json:"canonicalBaseUrl"`
		BapAdUrl         string `json:"bapAdUrl"`
		BapSearchUrl     string `json:"bapSearchUrl"`
		BapCollectionUrl string `json:"bapCollectionUrl"`
		BapBrowseUrl     string `json:"bapBrowseUrl"`
		ImageCdnUrl      string `json:"imageCdnUrl"`
		AutoCompleteApi  string `json:"autoCompleteApi"`
		FavoritesApi     string `json:"favoritesApi"`
		MapApi           string `json:"mapApi"`
		NavTiming        struct {
			Percentage int    `json:"percentage"`
			Url        string `json:"url"`
		} `json:"navTiming"`
		FeatureToggleMapping struct {
			MultiVariantJobRecommendations string `json:"multiVariantJobRecommendations"`
			MultiVariantJobRerank          string `json:"multiVariantJobRerank"`
			RealestateLettingsMedium       string `json:"realestateLettingsMedium"`
			ShowHeading                    string `json:"showHeading"`
			TjtHotjarEnabled               string `json:"tjtHotjarEnabled"`
		} `json:"featureToggleMapping"`
		ApplicationName string `json:"applicationName"`
	} `json:"runtimeConfig"`
	IsFallback   bool          `json:"isFallback"`
	Gssp         bool          `json:"gssp"`
	CustomServer bool          `json:"customServer"`
	AppGip       bool          `json:"appGip"`
	ScriptLoader []interface{} `json:"scriptLoader"`
}
