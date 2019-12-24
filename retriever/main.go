package main

import (
	"awesomeProject/retriever/mock"
	"awesomeProject/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url  = "http://www.imooc.com"

//组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string{
	s.Post(url, map[string]string {
		"contents": "another fake",
	})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url,
		map[string]string {
			"name": "ccmouse",
			"course": "golang",
		})
}

func main() {
	var r Retriever

	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(r)

	//Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}

	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever)  {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
		case *mock.Retriever:
			fmt.Println("Contents:", v.Contents)
		case *real.Retriever:
			fmt.Println("Useragent:", v.UserAgent)
	}
}