package hacknews

import (
	"testing"
	"github.com/PaulRosset/go-hacknews"
)

// Inside the test/ folder
// $ go test

func TestGetPostTopStories(t *testing.T) {
	init := hacknews.Initializer{"topstories", 40}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostTopStories] Expected error to be nil, got : ", err)
	}
	
	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostTopStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostTopStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}
	for _, post := range posts {
		if post.Id == 0 {
			t.Error("[TestGetPostTopStories] Id is mandatory, got : ", post.Id)
		}
	}
}

func TestGetPostNewStories(t *testing.T) {
	init := hacknews.Initializer{"newstories", 30}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostNewStories] Expected error to be nil, got : ", err)
	}
	
	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostNewStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostNewStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}

	for _, post := range posts {
		if post.Id == 0 {
			t.Error("[TestGetPostNewStories] Id is mandatory, got : ", post.Id)
		}
	}
}

func TestGetPostBestStories(t *testing.T) {
	init := hacknews.Initializer{"beststories", 20}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostBestStories] Expected error to be nil, got : ", err)
	}
	
	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostBestStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostBestStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}
	
	for _, post := range posts {
		if post.Id == 0 {
			t.Error("[TestGetPostBestStories] Id is mandatory, got : ", post.Id)
		}
	}
}

func TestGetPostAskStories(t *testing.T) {
	init := hacknews.Initializer{"askstories", 15}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostAskStories] Expected error to be nil, got : ", err)
	}

	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostAskStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostAskStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}

	for _, post := range posts {
		if post.Id == 0 {
			t.Error("TestGetPostAskStories] Id is mandatory, got : ", post.Id)
		}
	}
}

func TestGetPostShowStories(t *testing.T) {
	init := hacknews.Initializer{"showstories", 10}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostShowStories] Expected error to be nil, got : ", err)
	}

	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostShowStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostShowStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}

	for _, post := range posts {
		if post.Id == 0 {
			t.Error("TestGetPostShowStories] Id is mandatory, got : ", post.Id)
		}
	}
}

func TestGetPostJobStories(t *testing.T) {
	init := hacknews.Initializer{"jobstories", 5}
	codes, err := init.GetCodesStory()
	
	if err != nil {
		t.Error("[TestGetPostJobStories] Expected error to be nil, got : ", err)
	}

	posts, errPost := init.GetPostStory(codes)
	
	if errPost != nil {
		t.Error("[TestGetPostJobStories] Expected error to be nil, got : ", err)
	} 
	
	if posts == nil || len(posts) == 0 || len(posts) != init.NbPosts || cap(posts) == 0 {
		t.Error("[TestGetPostJobStories] Expected posts to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", posts, len(posts), cap(posts))
	}
	
	for _, post := range posts {
		if post.Id == 0 {
			t.Error("TestGetPostJobStories] Id is mandatory, got : ", post.Id)
		}
	}
}
