package hacknews

import (
	"github.com/PaulRosset/go-hacknews"
	"testing"
)

// Inside the test/ folder
// $ go test

func TestGetCodesTopStories(t *testing.T) {
	init := hacknews.Initializer{"topstories", 40}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodesTopStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodesTopStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}

func TestGetCodeNewStories(t *testing.T) {
	init := hacknews.Initializer{"newstories", 30}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodeNewStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodeNewStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}

func TestGetCodesBestStories(t *testing.T) {
	init := hacknews.Initializer{"beststories", 20}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodesBestStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodesBestStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}

func TestGetCodesAskStories(t *testing.T) {
	init := hacknews.Initializer{"askstories", 15}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodesAskStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodesAskStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}

func TestGetCodesShowStories(t *testing.T) {
	init := hacknews.Initializer{"showstories", 10}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodesShowStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodesShowStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}

func TestGetCodesJobStories(t *testing.T) {
	init := hacknews.Initializer{"jobstories", 5}
	codes, err := init.GetCodesStory()

	if err != nil {
		t.Error("[TestGetCodesJobStories] Expected error to be nil, got : ", err)
	}

	if codes == nil || len(codes) == 0 || len(codes) <= init.NbPosts || cap(codes) == 0 {
		t.Error("[TestGetCodesJobStories] Expected codes to be not nil or have a length > 0 and atleast > to the number of posts asked or capacity > 0, got : ", codes, len(codes), cap(codes))
	}
}
