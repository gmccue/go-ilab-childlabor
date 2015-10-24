package laborstats

import "encoding/json"

// Request path for Suggested Action data.
const suggestedActionAreaURI = "childlabor_actionarea"

type suggestedActionAreaAPI laborStatsAPI

type suggestedActionArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *suggestedActionAreaAPI) sendRequest() error {
	api.endpoint = buildEndpoint(suggestedActionAreaURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *suggestedActionAreaAPI) unmarshalData() ([]suggestedActionArea, error) {
	var suggestedActionAreas []suggestedActionArea

	err := json.Unmarshal(api.RawResponse, &suggestedActionAreas)
	if err != nil {
		return nil, err
	}

	return suggestedActionAreas, nil
}
