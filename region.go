package laborstats

import "encoding/json"

// Request path for Region data.
const regionURI = "childlabor_reg"

type RegionAPI LaborStatsAPI

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *RegionAPI) sendRequest() error {
	api.endpoint = buildEndpoint(regionURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *RegionAPI) unmarshalData() ([]Region, error) {
	var regions []Region

	err := json.Unmarshal(api.RawResponse, &regions)
	if err != nil {
		return nil, err
	}

	return regions, nil
}
