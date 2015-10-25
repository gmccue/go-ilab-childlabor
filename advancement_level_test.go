package laborstats

import "testing"

func TestAdvLvlUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/advancement_level.json")
	if err != nil {
		t.Error(err)
	}

	api := AdvancementLevelAPI{}
	api.RawResponse = dataMock

	result, err := api.unmarshalData()
	if err != nil {
		t.Error(err)
	}

	fRes := result[0]
	if fRes.ID != 1 {
		t.Error("Invalid ID for first result.")
	}

	lRes := result[2]
	if lRes.Name != "Significant Advancement" {
		t.Error("Invalid Name for last result.")
	}
}
