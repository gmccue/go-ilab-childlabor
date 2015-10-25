package laborstats

import "testing"

func TestCountryGoodsUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/country_goods.json")
	if err != nil {
		t.Error(err)
	}

	api := CountryGoodsAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	fRes := result[0]
	if fRes.CountryProfileID != 1 {
		t.Error("Invalid CountryProfileID: ", fRes.CountryProfileID)
	}
	if fRes.GoodID != 1 {
		t.Error("Invalid GoodID: ", fRes.GoodID)
	}
	if fRes.ChildLabor != false {
		t.Error("Invalid ChildLabor value: ", fRes.ChildLabor)
	}
	if fRes.ForcedLabor != false {
		t.Error("Invalid ForcedLabor value: ", fRes.ForcedLabor)
	}
	if fRes.ForcedChildLabor != false {
		t.Error("Invalid ForcedChildLabor value: ", fRes.ForcedChildLabor)
	}
}
