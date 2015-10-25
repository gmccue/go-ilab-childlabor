package laborstats

import "encoding/json"

// Request path for Country Statistics data.
const countryStatsURI = "childlabor_sta"

type CountryStatsAPI LaborStatsAPI

type CountryStat struct {
	CountryProfileID  int     `json:"country_profile_id"`
	CWAgeRange        string  `json:"cws_age_range,omitempty"`
	CWPercent         float64 `json:"cws_total_percentage_of_workin,omitempty"`
	CWPopulation      int     `json:"cws_total_working_population"`
	CWAgriculture     float64 `json:"cws_agriculture,omitempty"`
	CWService         float64 `json:"cws_services,omitempty"`
	CWIndustry        float64 `json:"cws_industry,omitempty"`
	SchoolAttYear     string  `json:"esas_year,omitempty"`
	SchoolAttAgeRange string  `json:"esas_age_range,omitempty"`
	SchoolAttPercent  float64 `json:"esas_percentage,omitempty"`
	CWASYear          string  `json:"cwas_year,omitempty"`
	CWASAgeRange      string  `json:"cwas_age_range,omitempty"`
	CWASTotal         float64 `json:"cwas_total,omitempty"`
	PCRYear           string  `json:"upcr_year,omitempty"`
	PCRRate           float64 `json:"upcr_rate,omitempty"`
}

func (api *CountryStatsAPI) sendRequest() error {
	api.endpoint = buildEndpoint(countryStatsURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *CountryStatsAPI) unmarshalData() ([]CountryStat, error) {
	var countryStats []CountryStat

	err := json.Unmarshal(api.RawResponse, &countryStats)
	if err != nil {
		return nil, err
	}

	return countryStats, nil
}
