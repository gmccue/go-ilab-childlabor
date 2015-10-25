package laborstats

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const testAPIKey = "xx"

func getDataMock(mockPath string) ([]byte, error) {
	dataMock, err := ioutil.ReadFile(mockPath)
	if err != nil {
		return nil, err
	}

	return dataMock, nil
}

func TestAddFilter(t *testing.T) {
	a := LaborStatsAPI{}
	testPath := "myPath"

	a.AddFilter("limit", "10")

	endpoint := buildEndpoint(testPath, a.Filters)

	if endpoint.String() != fmt.Sprintf("%s://%s/%s/%s/%s/%s", apiScheme, apiHost, apiPath, testPath, "limit", "10") {
		t.Error("Invalid endpoint built: ", endpoint.String())
	}
}

func TestFilterIsValid(t *testing.T) {
	invalidFilter := "not_a_filter"

	if filterIsValid(invalidFilter) {
		t.Error("Invalid filter detected as valid: ", invalidFilter)
	}
}

func TestBuildEndpoint(t *testing.T) {
	testPath := "myPath"
	testFilters := map[string]string{}
	endpoint := buildEndpoint(testPath, testFilters)

	if endpoint.String() != fmt.Sprintf("%s://%s/%s/%s", apiScheme, apiHost, apiPath, testPath) {
		t.Error("Invalid endpoint built: ", endpoint.String())
	}
}
