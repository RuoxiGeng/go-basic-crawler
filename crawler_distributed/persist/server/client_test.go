package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"awesomeProject/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
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

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
