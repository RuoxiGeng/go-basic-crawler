package main

import (
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
