// Package laborstats provides programmatic access to the ILAB Sweat & Toil
// API (http://developer.dol.gov/others/sweat-and-toil/).
package laborstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiScheme = "https"
	apiHost   = "data.dol.gov"
	apiPath   = "get"

	// Custom request header containing the API secret key
	secretKeyHeader = "X-API-KEY"
)

var (
	LaborStatsAPIError   = errors.New("The API request returned an error.")
	invalidFilterError   = errors.New("Invalid query parameter provided.")
	invalidResponseError = errors.New("The HTTP request failed.")

	// validFiltersKeys holds an array of currently available query filters.
	validFilterKeys = []string{"limit", "date_column", "start_date", "end_date", "order"}
)

type QueryFilters map[string]string

// APIError holds error information returned from an API request.
type APIError struct {
	status  bool   `json:"status"`
	Message string `json:"error"`
}

type LaborStatsAPI struct {
	Debug       bool
	Filters     QueryFilters
	RawResponse []byte
	SecretKey   string
	endpoint    *url.URL
}

type QueryRunner interface {
	sendRequest() error
	unmarshalData() error
}

// lsbool is a custom boolean type for unmarshaling JSON
type lsbool bool

// NewLaborStatsAPI configures and returns a new API instance.
func NewLaborStatsAPI(secretKey string) *LaborStatsAPI {
	return &LaborStatsAPI{
		Filters:   make(QueryFilters),
		SecretKey: secretKey,
	}
}

// QueryAdvancementLevel submits an API request against the AdvancementLevel
// endpoint.
func (api *LaborStatsAPI) QueryAdvancementLevel() ([]AdvancementLevel, error) {
	a := AdvancementLevelAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryCountry submits an API request against the Country endpoint.
func (api *LaborStatsAPI) QueryCountry() ([]Country, error) {
	a := CountryAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryCountryGoods submits an API request against the Country Goods endpoint.
func (api *LaborStatsAPI) QueryCountryGoods() ([]CountryGood, error) {
	a := CountryGoodsAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// QueryCountryProfile submits an API request against the Country Profile
// endpoint.
func (api *LaborStatsAPI) QueryCountryProfile() ([]CountryProfile, error) {
	a := CountryProfileAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// QueryCountryStats submits an API request against the Country Statistics
// endpoint.
func (api *LaborStatsAPI) QueryCountryStats() ([]CountryStat, error) {
	a := CountryStatsAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryGood submits an API request against the "Good" endpoint.
func (api *LaborStatsAPI) QueryGood() ([]Good, error) {
	a := GoodAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// QueryRegion submits an API request against the Region endpoint.
func (api *LaborStatsAPI) QueryRegion() ([]Region, error) {
	a := RegionAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// QuerySector submits an API request against the Sector endpoint.
func (api *LaborStatsAPI) QuerySector() ([]Sector, error) {
	a := SectorAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// QuerySuggestedActionArea submits an API request against the Suggested Action
// Area endpoint.
func (api *LaborStatsAPI) QuerySuggestedActionArea() ([]SuggestedActionArea, error) {
	a := SuggestedActionAreaAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil

}

// QuerySuggestedActions submits an API request against the Suggested Actions
// endpoint.
func (api *LaborStatsAPI) QuerySuggestedActions() ([]SuggestedAction, error) {
	a := SuggestedActionAPI{
		Debug:     api.Debug,
		SecretKey: api.SecretKey,
	}

	err := a.sendRequest()
	if err != nil {
		return nil, err
	}

	res, err := a.unmarshalData()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AddFilter adds a filter parameter to the API request.
// The currently available filters are "limit", "start_date", "end_date",
// and "order".
func (api *LaborStatsAPI) AddFilter(filterName string, filterValue string) error {
	if !filterIsValid(filterName) {
		return invalidFilterError
	}

	if len(api.Filters) == 0 {
		api.Filters = make(QueryFilters)
		api.Filters[filterName] = filterValue
	} else {
		api.Filters[filterName] = filterValue
	}

	return nil
}

func filterIsValid(filterKey string) bool {
	for _, x := range validFilterKeys {
		if x == filterKey {
			return true
		}
	}

	return false
}

func buildEndpoint(path string, filterMap QueryFilters) *url.URL {
	var filters []string

	for key, val := range filterMap {
		filters = append(filters, fmt.Sprintf("%s/%s", key, val))
	}

	queryPath := fmt.Sprintf("%s/%s", apiPath, path)

	if len(filters) > 0 {
		queryPath = fmt.Sprintf("%s/%s/%s", apiPath, path, strings.Join(filters, "/"))
	}

	url := &url.URL{
		Scheme: apiScheme,
		Host:   apiHost,
		Path:   queryPath,
	}

	return url
}

func doRequest(endpointURL string, secretKey string, debug bool) ([]byte, error) {
	if debug {
		log.Printf("API endpoint URL: %s", endpointURL)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(secretKeyHeader, secretKey)

	if debug {
		log.Printf("HTTP request headers: %v", req.Header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if debug {
		log.Printf("Response body: %v", string(body))
	}

	apiErr := unmarshalErrorResponse(body)
	if apiErr != nil {
		return nil, apiErr
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s HTTP status code returned was: %d.", invalidResponseError, resp.StatusCode)
	}

	return body, nil
}

// unmarshalErrorResponse attempts to unmarshal an API error message if one
// exists. If an error is found, the details are returned.
func unmarshalErrorResponse(b []byte) error {
	apiErr := APIError{}

	err := json.Unmarshal(b, &apiErr)
	if err == nil {
		return fmt.Errorf("%s The error message was: %+s", LaborStatsAPIError, apiErr.Message)
	}

	return nil
}

// UnmarshalJSON is a custom implementation of the UnmarshalJSON interface for
// values returned from the Labor Stats API. The JSON returned by the API to
// represent boolean values is 0 or 1.
func (bv *lsbool) UnmarshalJSON(b []byte) error {
	boolVal := string(b) == "1"

	*bv = lsbool(boolVal)

	return nil
}
