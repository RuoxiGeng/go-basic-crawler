package main

import (
	"awesomeProject/crawler_distributed/persist"
	"awesomeProject/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
