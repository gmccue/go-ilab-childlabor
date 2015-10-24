package laborstats

import "encoding/json"

// Request path for Advancement Level data.
const advancementLevelURI = "childlabor_advlvl"

type advancementLevelAPI laborStatsAPI

type advancementLevel struct {
	ID   int    `json:"id"`
	Name string `json:"advancement_name"`
}

func (api *advancementLevelAPI) sendRequest() error {
	api.endpoint = buildEndpoint(advancementLevelURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *advancementLevelAPI) unmarshalData() ([]advancementLevel, error) {
	var advLvl []advancementLevel

	err := json.Unmarshal(api.RawResponse, &advLvl)

	if err != nil {
		return nil, err
	}

	return advLvl, nil
}
