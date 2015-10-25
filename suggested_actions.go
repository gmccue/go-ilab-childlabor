package laborstats

import "encoding/json"

// Request path for Suggested Action data.
const suggestedActionURI = "childlabor_action"

type SuggestedActionAPI LaborStatsAPI

type SuggestedAction struct {
	ID               int    `json:"id"`
	CountryProfileID int    `json:"country_profile_id"`
	ActionAreaID     int    `json:"area_id"`
	Name             string `json:"name,omitempty"`
	Year             string `json:"year,omitempty"`
}

func (api *SuggestedActionAPI) sendRequest() error {
	api.endpoint = buildEndpoint(suggestedActionURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *SuggestedActionAPI) unmarshalData() ([]SuggestedAction, error) {
	var suggestedActions []SuggestedAction

	err := json.Unmarshal(api.RawResponse, &suggestedActions)
	if err != nil {
		return nil, err
	}

	return suggestedActions, nil
}
