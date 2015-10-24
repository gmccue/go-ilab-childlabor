package laborstats

import "encoding/json"

// Request path for Country Goods data.
const countryGoodsURI = "childlabor_cty_goo"

type countryGoodsAPI laborStatsAPI

type countryGood struct {
	CountryProfileID int    `json:"country_profile_id,omitempty"`
	GoodID           int    `json:"good_id,omitempty"`
	ChildLabor       lsbool `json:"child_Labor,omitempty"`
	ForcedLabor      lsbool `json:"forced_labor,omitempty"`
	ForcedChildLabor lsbool `json:"forced_child_labor,omitempty"`
}

func (api *countryGoodsAPI) sendRequest() error {
	api.endpoint = buildEndpoint(countryGoodsURI, api.filters)

	rawResponse, err := doRequest(api.endpoint.String(), api.SecretKey, api.Debug)
	if err != nil {
		return err
	}

	api.RawResponse = rawResponse

	return nil
}

func (api *countryGoodsAPI) unmarshalData() ([]countryGood, error) {
	var countryGoods []countryGood

	err := json.Unmarshal(api.RawResponse, &countryGoods)
	if err != nil {
		return nil, err
	}

	return countryGoods, nil
}
