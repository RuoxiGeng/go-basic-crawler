package main

import (
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://www.7799520.com/user/2744596.html",
		Parser: worker.SerializedParser{
			Name: "ParseProfile",
			Args: "水蜜桃向顺利",
		},

	}
	var result worker.ParseResult
	err = client.Call("CrawlService.Process", req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
