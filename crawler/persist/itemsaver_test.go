package persist

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
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

	client , err := elastic.NewClient(
		//Must turn off sniff in dockerTy
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", expected, actual)
	}
}