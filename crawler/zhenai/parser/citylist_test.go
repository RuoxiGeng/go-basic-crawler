package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	expectedUrls := []string {
		"http://www.7799520.com/zhenghun/anhui",
		"http://www.7799520.com/zhenghun/aomen",
		"http://www.7799520.com/zhenghun/anqing",
	}

	const resultSize = 389

	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d " +
			"requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but " +
				"was %s", i, url, result.Requests[i].Url)
		}
	}
}
