package laborstats

import "encoding/json"

// Request path for Sector data.
const sectorURI = "childlabor_sec"

type sectorAPI laborStatsAPI

type sector struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *sectorAPI) sendRequest() error {
	api.endpoint = buildEndpoint(sectorURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *sectorAPI) unmarshalData() ([]sector, error) {
	var sectors []sector

	err := json.Unmarshal(api.RawResponse, &sectors)
	if err != nil {
		return nil, err
	}

	return sectors, nil
}
