# go-hacknews
Tiny SDK for HackerNews API written in go.

[Official Hackernews API](https://github.com/HackerNews/API)

## Install

```
 go get github.com/PaulRosset/go-hacknews
```

## Usage

Few examples are available inside the **examples/** folder.

```go
package main

import (
	    "fmt"
	    "github.com/PaulRosset/go-hacknews"
)

func main() {

	// Init struct with the kind of story you want
	// (topstories/newstories/beststories/askstories/showstories/jobstories)
	// and the number of posts that you want to fetch.
	init := hacknews.Initializer{"topstories", 10}

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
	// Here, get Title and Url of post.
	// Note : Field are empty, if no data belong to them
	for _, post := range posts {
		fmt.Printf("Title : %v // Url : %v\n", post.Title ,post.Url)
	}

}



```

## Test

Tests are available :

```
$ cd test/
$ go test
```

## License

MIT
