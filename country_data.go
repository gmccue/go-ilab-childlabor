package laborstats

import "encoding/json"

// Request path for Country Data data.
const countryDataURI = "childlabor_mas"

type CountryDataAPI LaborStatsAPI

type CountryData struct {
	CountryProfileID          int    `json:"country_profile_id"`
	C138Ratified              string `json:"c_138_ratified,omitempty"`
	C182Ratified              string `json:"c_182_ratified,omitempty"`
	CRCRatificationStatus     string `json:"convention_on_the_rights_of_th,omitempty"`
	CRCCSARatificationStatus  string `json:"crc_commercial_sexual_exploita,omitempty"`
	CRCACRatificationStatus   string `json:"crc_armed_conflict_ratified,omitempty"`
	PalermoRatificationStatus string `json:"palermo_ratified,omitempty"`
	MinWorkAgeStatus          string `json:"minimum_age_for_work_establish,omitempty"`
	MinWorkAge                string `json:"minimum_age_for_work,omitempty"`
	MinHazWorkAgeStatus       string `json:"minimum_age_for_hazardous_work_established,omitempty"`
	MinHazWorkAge             string `json:"minimum_age_for_hazardous_work,omitempty"`
	CompEdAgeStatus           string `json:"compulsory_education_age_estab,omitempty"`
	CompEdAge                 string `json:"minimum_age_for_compulsory_edu",omitempty`
	FreePubEdStatus           string `json:"free_public_education_establis,omitepty"`
}

func (api *CountryDataAPI) sendRequest() error {
	api.endpoint = buildEndpoint(countryDataURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *CountryDataAPI) unmarshalData() ([]CountryData, error) {
	var countryData []CountryData

	err := json.Unmarshal(api.RawResponse, &countryData)
	if err != nil {
		return nil, err
	}

	return countryData, nil
}
