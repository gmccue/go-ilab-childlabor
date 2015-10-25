package laborstats

import "testing"

func TestCountryStatsUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/country_stats.json")
	if err != nil {
		t.Error(err)
	}

	api := CountryStatsAPI{}
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
	if fRes.CWAgeRange != "5-14" {
		t.Error("Invalid CWAgeRange value: ", fRes.CWAgeRange)
	}
	if fRes.CWPercent != 7.5 {
		t.Error("Invalid CWPercent value: ", fRes.CWPercent)
	}
	if fRes.CWPopulation != 673949 {
		t.Error("Invalid CWPopulation value: ", fRes.CWPopulation)
	}
	if fRes.CWAgriculture != 0 {
		t.Error("Invalid CWAgriculture value: ", fRes.CWAgriculture)
	}
	if fRes.CWService != 0 {
		t.Error("Invalid CWService value: ", fRes.CWService)
	}
	if fRes.CWIndustry != 0 {
		t.Error("Invalid CWIndustry value: ", fRes.CWIndustry)
	}
	if fRes.SchoolAttYear != "2010-11" {
		t.Error("Invalid ScoolAttYear value: ", fRes.SchoolAttYear)
	}
	if fRes.SchoolAttAgeRange != "5-14" {
		t.Error("Invalid SchoolAttAgeRange value: ", fRes.SchoolAttAgeRange)
	}
	if fRes.SchoolAttPercent != 41.8 {
		t.Error("Invalid SchoolAttPercent value: ", fRes.SchoolAttPercent)
	}
	if fRes.CWASYear != "2010-11" {
		t.Error("Invalid CWASYear value: ", fRes.CWASYear)
	}
	if fRes.CWASAgeRange != "7-14" {
		t.Error("Invalid CWASAgeRange value: ", fRes.CWASAgeRange)
	}
	if fRes.CWASTotal != 4.6 {
		t.Error("Invalid CWASTotal value: ", fRes.CWASTotal)
	}
	if fRes.PCRYear != "0000" {
		t.Error("Invalid PCRYear value: ", fRes.PCRYear)
	}
	if fRes.PCRRate != 0 {
		t.Error("Invalid PCRRate value: ", fRes.PCRRate)
	}
}
