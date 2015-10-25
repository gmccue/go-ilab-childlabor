package laborstats

import "testing"

func TestSuggestedActionsUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/suggested_actions.json")
	if err != nil {
		t.Error(err)
	}

	api := SuggestedActionAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	if len(result) != 2 {
		t.Error("Invalid result length: ", len(result))
	}

	fRes := result[0]
	if fRes.ID != 1 {
		t.Error("Invalid ID value: ", fRes.ID)
	}
	if fRes.CountryProfileID != 1 {
		t.Error("Invalid CountryProfileID value: ", fRes.CountryProfileID)
	}
	if fRes.ActionAreaID != 1 {
		t.Error("Invalid ActionAreaID value: ", fRes.ActionAreaID)
	}
	if fRes.Name != "Create better laws" {
		t.Error("Invalid Name value: ", fRes.Name)
	}
	if fRes.Year != "2013 - 2014" {
		t.Error("Invalid Year value: ", fRes.Year)
	}
}
