package main

import (
	"fmt"
	"github.com/PaulRosset/go-hacknews"
)

// go run examples/JobStoryTitle/main.go

func main() {
	// Init struct with the kind of story you want (topstories/newstories/beststories/askstories/showstories/jobstories) and the number of posts do you want to fetch.
	init := hacknews.Initializer{"jobstories", 5}

	// Get the code of posts.
	// Return a slice of int with the entry id if everything is ok or return an error.
	codes, err := init.GetCodesStory()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the posts thanks their id fetched above.
	// Return a slice of Post type with a readable format in go or return an err if fail.
	posts, err := init.GetPostStory(codes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over the slice to get what we want.
	// Here, get Title of post.
	// Note : Field are empty, if no data belong to them
	for _, post := range posts {
		fmt.Printf("Title : %v\n", post.Title)
	}
}
