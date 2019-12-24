package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://www.7799520.com/user/2744596.html", "水蜜桃向顺利")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0]
	expected := engine.Item{
		Url : "http://www.7799520.com/user/2744596.html",
		Type: "zhenai",
		Id: "2744596",
		Payload: model.Profile{
			Age:        53,
			Name:       "水蜜桃向顺利",
			Gender:     "男",
			Height:     167,
			Income:     "5000-10000",
			Marriage:   "未婚",
			Education:  "高中及以下",
			Occupation: "回头告诉你",
			Hukou:      "未填写",
			Xingzuo:    "巨蟹座",
		},
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
