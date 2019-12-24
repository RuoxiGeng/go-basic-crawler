package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	itemsaver "awesomeProject/crawler_distributed/persist/client"
	"awesomeProject/crawler_distributed/rpcsupport"
	worker "awesomeProject/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://www.7799520.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		log.Printf("Connected to %s", h)
		if err == nil {
			clients = append(clients, client)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}