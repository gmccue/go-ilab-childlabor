package laborstats

import "encoding/json"

// Request path for Advancement Level data.
const advancementLevelURI = "childlabor_advlvl"

type AdvancementLevelAPI LaborStatsAPI

type AdvancementLevel struct {
	ID   int    `json:"id"`
	Name string `json:"advancement_name"`
}

func (api *AdvancementLevelAPI) sendRequest() error {
	api.endpoint = buildEndpoint(advancementLevelURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *AdvancementLevelAPI) unmarshalData() ([]AdvancementLevel, error) {
	var advLvl []AdvancementLevel

	err := json.Unmarshal(api.RawResponse, &advLvl)

	if err != nil {
		return nil, err
	}

	return advLvl, nil
}
