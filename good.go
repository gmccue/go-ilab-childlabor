package laborstats

import "encoding/json"

// Request path for "Good" data.
const goodURI = "childlabor_goo"

type goodAPI laborStatsAPI

type good struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SectorID int    `json:"sector_id,omitempty"`
}

func (api *goodAPI) sendRequest() error {
	api.endpoint = buildEndpoint(goodURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *goodAPI) unmarshalData() ([]good, error) {
	var goods []good

	err := json.Unmarshal(api.RawResponse, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
