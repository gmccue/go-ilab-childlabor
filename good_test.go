package laborstats

import "testing"

func TestGoodUnmarshalData(t *testing.T) {
	dataMock, err := getDataMock("./testdata/good.json")
	if err != nil {
		t.Error(err)
	}

	api := goodAPI{}
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
		t.Error("Invalid ID: ", fRes.ID)
	}
	if fRes.Name != "Bricks" {
		t.Error("Invalid Name: ", fRes.Name)
	}
	if fRes.SectorID != 1 {
		t.Error("Invalid SectorID value: ", fRes.SectorID)
	}
}
