package laborstats

import "testing"

func TestCountryDataUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/country_data.json")
	if err != nil {
		t.Error(err)
	}

	api := countryDataAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	if len(result) != 2 {
		t.Error("Invalid result length: ", len(result))
	}

	fRes := result[0]

	if fRes.CountryProfileID != 1 {
		t.Error("Invalid CountryProfileID value: ", fRes.CountryProfileID)
	}

	if fRes.C138Ratified != "Yes" {
		t.Error("Invalid C138Ratified value.")
	}

	if fRes.C182Ratified != "Yes" {
		t.Error("Invalid C182Ratified value.")
	}

	if fRes.CRCRatificationStatus != "Yes" {
		t.Error("Invalid CRCRatificationStatus value.")
	}

	if fRes.CRCCSARatificationStatus != "Yes" {
		t.Error("Invalid CRCCSARatificationStatus value.")
	}

	if fRes.CRCACRatificationStatus != "Yes" {
		t.Error("Invalid CRCACRatificationStatus value.")
	}

	if fRes.PalermoRatificationStatus != "Yes" {
		t.Error("Invalid PalermoRatificationStatus value.")
	}

	if fRes.MinWorkAgeStatus != "Yes" {
		t.Error("Invalid MinWorkAgeStatus value.")
	}

	if fRes.MinWorkAge != "15" {
		t.Error("Invalid MinWorkAge value.")
	}

	if fRes.MinHazWorkAgeStatus != "Yes" {
		t.Error("Invalid MinHazWorkAgeStatus value.")
	}

	if fRes.MinHazWorkAge != "18" {
		t.Error("Invalid MinHazWorkAge value.")
	}

	if fRes.CompEdAgeStatus != "Yes" {
		t.Error("Invalid CompEdAgeStatus value.")
	}

	if fRes.CompEdAge != "15" {
		t.Error("Invalid CompEdAge value.")
	}

	if fRes.FreePubEdStatus != "Yes" {
		t.Error("Invalid FreePubEdStatus value.")
	}
}
