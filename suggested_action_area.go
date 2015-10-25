package laborstats

import "encoding/json"

// Request path for Suggested Action data.
const suggestedActionAreaURI = "childlabor_actionarea"

type SuggestedActionAreaAPI LaborStatsAPI

type SuggestedActionArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *SuggestedActionAreaAPI) sendRequest() error {
	api.endpoint = buildEndpoint(suggestedActionAreaURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *SuggestedActionAreaAPI) unmarshalData() ([]SuggestedActionArea, error) {
	var suggestedActionAreas []SuggestedActionArea

	err := json.Unmarshal(api.RawResponse, &suggestedActionAreas)
	if err != nil {
		return nil, err
	}

	return suggestedActionAreas, nil
}
