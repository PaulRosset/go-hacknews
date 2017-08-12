# go-hacknews 
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badge/)
[![Travis](https://img.shields.io/travis/rust-lang/rust.svg)](https://travis-ci.org/PaulRosset/go-hacknews)
[![Go Report Card](https://goreportcard.com/badge/github.com/PaulRosset/go-hacknews)](https://goreportcard.com/report/github.com/PaulRosset/go-hacknews)
[![](https://cover.run/go/github.com/PaulRosset/go-hacknews.svg)](https://cover.run/go/github.com/PaulRosset/go-hacknews)
[![](http://godoc.org/github.com/PaulRosset/go-hacknews?status.svg)](http://godoc.org/github.com/PaulRosset/go-hacknews)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)


Tiny utility Go client for HackerNews API.

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
	// Here, get Title and Url of post. But we can access all fields from the official hackernews api
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
