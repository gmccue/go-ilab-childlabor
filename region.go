package laborstats

import "encoding/json"

// Request path for Region data.
const regionURI = "childlabor_reg"

type regionAPI laborStatsAPI

type region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *regionAPI) sendRequest() error {
	api.endpoint = buildEndpoint(regionURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *regionAPI) unmarshalData() ([]region, error) {
	var regions []region

	err := json.Unmarshal(api.RawResponse, &regions)
	if err != nil {
		return nil, err
	}

	return regions, nil
}
