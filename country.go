package laborstats

import "encoding/json"

// Request path for Country data.
const countryURI = "childlabor_cty"

type CountryAPI LaborStatsAPI

type Country struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RegionID int    `json:"region_id,omitempty"`
	ISO2     string `json:"iso2,omitempty"`
	ISO3     string `json:"iso3,omitempty"`
}

func (api *CountryAPI) sendRequest() error {
	api.endpoint = buildEndpoint(countryURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *CountryAPI) unmarshalData() ([]Country, error) {
	var country []Country

	err := json.Unmarshal(api.RawResponse, &country)
	if err != nil {
		return nil, err
	}

	return country, nil
}
