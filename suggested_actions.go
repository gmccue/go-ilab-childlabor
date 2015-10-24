package laborstats

import "encoding/json"

// Request path for Suggested Action data.
const suggestedActionURI = "childlabor_action"

type suggestedActionAPI laborStatsAPI

type suggestedAction struct {
	ID               int    `json:"id"`
	CountryProfileID int    `json:"country_profile_id"`
	ActionAreaID     int    `json:"area_id"`
	Name             string `json:"name,omitempty"`
	Year             string `json:"year,omitempty"`
}

func (api *suggestedActionAPI) sendRequest() error {
	api.endpoint = buildEndpoint(suggestedActionURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *suggestedActionAPI) unmarshalData() ([]suggestedAction, error) {
	var suggestedActions []suggestedAction

	err := json.Unmarshal(api.RawResponse, &suggestedActions)
	if err != nil {
		return nil, err
	}

	return suggestedActions, nil
}
