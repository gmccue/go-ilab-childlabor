package laborstats

import "testing"

func TestCountryProfileUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/country_profile.json")
	if err != nil {
		t.Error(err)
	}

	api := countryProfileAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	fRes := result[0]
	if fRes.ID != 1 {
		t.Error("Invalid ID: ", fRes.ID)
	}
	if fRes.CountryID != 1 {
		t.Error("Invalid CountryID: ", fRes.CountryID)
	}
	if fRes.ProfileYear != 2014 {
		t.Error("Invalid ProfileYear: ", fRes.ProfileYear)
	}
	if fRes.AdLevelID != 1 {
		t.Error("Invalid AdLevelID: ", fRes.AdLevelID)
	}
	if fRes.Description != "A general description of a country profile." {
		t.Error("Invalid Description: ", fRes.CountryID)
	}
}
