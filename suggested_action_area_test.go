package laborstats

import "testing"

func TestSuggestedActionAreaUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/suggested_action_area.json")
	if err != nil {
		t.Error(err)
	}

	api := SuggestedActionAreaAPI{}
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
	if fRes.Name != "Legal Framework" {
		t.Error("Invalid Name value: ", fRes.Name)
	}
}
