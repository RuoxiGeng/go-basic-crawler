package main

import (
	"fmt"
	"regexp"
)

const text  =  `
My email is geng@gmail.com
My email is geng1@gmail.com
My email is geng2@gmail.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
