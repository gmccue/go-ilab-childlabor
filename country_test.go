package laborstats

import "testing"

func TestCountryUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/country.json")
	if err != nil {
		t.Error(err)
	}

	api := countryAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	fRes := result[0]
	if fRes.Name != "Country One" {
		t.Error("Invalid Name for first result.")
	}

	if fRes.RegionID != 1 {
		t.Error("Invalid ID for first result.")
	}

	if fRes.ISO2 != "C1" {
		t.Error("Invalid ISO2 for first result.")
	}

	if fRes.ISO3 != "CT1" {
		t.Error("Invalid ISO3 for first result.")
	}

	lRes := result[1]
	if lRes.Name != "Country Two" {
		t.Error("Invalid Name for last result.")
	}

	if lRes.RegionID != 2 {
		t.Error("Invalid RegionID for last result.")
	}

	if lRes.ISO2 != "C2" {
		t.Error("Invalid ISO2 for last result.")
	}

	if lRes.ISO3 != "CT2" {
		t.Error("Invalid ISO3 for last result.")
	}
}
