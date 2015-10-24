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

	// Request path to designate a date filter.
	dateFilterIndicator = "date_column"

	// Request path to designate a limit filter.
	limitFilterIndicator = "limit"

	// Request path to designate an order filter
	orderFilterIndicator = "orderby"

	// Custom request header containing the API secret key
	secretKeyHeader = "X-API-KEY"
)

var (
	laborStatsAPIError   = errors.New("The API request returned an error.")
	invalidFilterError   = errors.New("Invalid query parameter provided.")
	invalidResponseError = errors.New("The HTTP request failed.")

	// validFiltersKeys holds an array of currently available query filters.
	validFilterKeys = []string{"limit", "start_date", "end_date", "order"}
)

type queryFilters map[string]string

// APIError holds error information returned from an API request.
type APIError struct {
	status  bool   `json:"status"`
	Message string `json:"error"`
}

type laborStatsAPI struct {
	Debug       bool
	RawResponse []byte
	SecretKey   string
	endpoint    *url.URL
	filters     queryFilters
}

type queryRunner interface {
	sendRequest() error
	unmarshalData() error
}

// lsbool is a custom boolean type for unmarshaling JSON
type lsbool bool

// NewLaborStatsAPI configures and returns a new API instance.
func NewLaborStatsAPI(secretKey string) *laborStatsAPI {
	return &laborStatsAPI{
		filters:   make(queryFilters),
		SecretKey: secretKey,
	}
}

// QueryAdvancementLevel submits an API request against the AdvancementLevel
// endpoint.
func (api *laborStatsAPI) QueryAdvancementLevel() ([]advancementLevel, error) {
	a := advancementLevelAPI{
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
func (api *laborStatsAPI) QueryCountry() ([]country, error) {
	a := countryAPI{
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
func (api *laborStatsAPI) QueryCountryGoods() ([]countryGood, error) {
	a := countryGoodsAPI{
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
func (api *laborStatsAPI) QueryCountryProfile() ([]countryProfile, error) {
	a := countryProfileAPI{
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
func (api *laborStatsAPI) QueryCountryStats() ([]countryStat, error) {
	a := countryStatsAPI{
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
func (api *laborStatsAPI) QueryGood() ([]good, error) {
	a := goodAPI{
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
func (api *laborStatsAPI) QueryRegion() ([]region, error) {
	a := regionAPI{
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
func (api *laborStatsAPI) QuerySector() ([]sector, error) {
	a := sectorAPI{
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
func (api *laborStatsAPI) QuerySuggestedActionArea() ([]suggestedActionArea, error) {
	a := suggestedActionAreaAPI{
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
func (api *laborStatsAPI) QuerySuggestedActions() ([]suggestedAction, error) {
	a := suggestedActionAPI{
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
func (api *laborStatsAPI) AddFilter(filterName string, filterValue string) error {
	if !filterIsValid(filterName) {
		return invalidFilterError
	}

	if len(api.filters) == 0 {
		api.filters = make(queryFilters)
		api.filters[filterName] = filterValue
	} else {
		api.filters[filterName] = filterValue
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

func buildEndpoint(path string, filterMap queryFilters) *url.URL {
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
		return fmt.Errorf("%s The error message was: %+s", laborStatsAPIError, apiErr.Message)
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
