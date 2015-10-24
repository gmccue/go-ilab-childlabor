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
	validFilterKeys      = []string{"limit", "start_date", "end_date", "order"}
)

type queryFilters map[string]string

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

type queryBuilder interface {
	sendRequest() error
	unmarshalData() error
}

// lsbool is a custom boolean type for unmarshaling JSON
type lsbool bool

func NewLaborStatsAPI(secretKey string) *laborStatsAPI {
	return &laborStatsAPI{
		filters:   make(queryFilters),
		SecretKey: secretKey,
	}
}

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
