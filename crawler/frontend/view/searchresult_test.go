package view

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/frontend/model"
	common "awesomeProject/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url : "http://www.7799520.com/user/2744596.html",
		Type: "zhenai",
		Id: "2744596",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
