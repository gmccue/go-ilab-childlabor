package laborstats

import "encoding/json"

// Request path for Sector data.
const sectorURI = "childlabor_sec"

type SectorAPI LaborStatsAPI

type Sector struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *SectorAPI) sendRequest() error {
	api.endpoint = buildEndpoint(sectorURI, api.Filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *SectorAPI) unmarshalData() ([]Sector, error) {
	var sectors []Sector

	err := json.Unmarshal(api.RawResponse, &sectors)
	if err != nil {
		return nil, err
	}

	return sectors, nil
}
