package main

import (
	"fmt"
	"../"
)

func main() {
	init := hacknews.Initializer{"topstories", 10}
	codes, err := init.GetCodesStory()
	if err != nil {
		fmt.Println(err)
		return
	}
	post, err := init.GetPostStory(codes)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range post {
		fmt.Printf("%v\n", v.Url)
	}
}


