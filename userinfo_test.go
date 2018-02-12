package igmedia

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetUserInfoNoLogin() {
	user, err := GetUserInfoNoLogin("instagram")
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Id)
	fmt.Println(user.Biography)
	// Output:
	// 25025320
	// Discovering — and telling — stories from around the world. Curated by Instagram’s community team.
}

func ExampleGetUserId() {
	fmt.Println(GetUserId("instagram"))
	// Output: 25025320 <nil>
}

func ExampleGetAllPostCode(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	codes, err := mgr.GetAllPostCode(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}
