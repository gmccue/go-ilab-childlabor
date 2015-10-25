package laborstats

import "encoding/json"

// Request path for Country Profile data.
const countryProfileURI = "childlabor_pro"

type CountryProfileAPI LaborStatsAPI

type CountryProfile struct {
	ID          int    `json:"id"`
	CountryID   int    `json:"country_id"`
	ProfileYear int    `json:"profile_year,omitempty"`
	AdLevelID   int    `json:"advancement_id,omitempty"`
	Description string `json:"description,omitempty"`
}

func (api *CountryProfileAPI) sendRequest() error {
	api.endpoint = buildEndpoint(countryProfileURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *CountryProfileAPI) unmarshalData() ([]CountryProfile, error) {
	var countryProfiles []CountryProfile

	err := json.Unmarshal(api.RawResponse, &countryProfiles)
	if err != nil {
		return nil, err
	}

	return countryProfiles, nil
}
