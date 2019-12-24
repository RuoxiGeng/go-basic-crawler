package client

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializseRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan

		err := c.Call("CrawlService.Process", sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}
}