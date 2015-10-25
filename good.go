package laborstats

import "encoding/json"

// Request path for "Good" data.
const goodURI = "childlabor_goo"

type GoodAPI LaborStatsAPI

type Good struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SectorID int    `json:"sector_id,omitempty"`
}

func (api *GoodAPI) sendRequest() error {
	api.endpoint = buildEndpoint(goodURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *GoodAPI) unmarshalData() ([]Good, error) {
	var goods []Good

	err := json.Unmarshal(api.RawResponse, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
