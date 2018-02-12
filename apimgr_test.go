package igmedia

import (
	"os"
	"testing"
)

func TestIGApiManager(t *testing.T) {
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
		sm, err := mgr.GetPostInfo(code)
		if err != nil {
			t.Error(err)
			return
		}
		//JsonPrettyPrint(sm)
		printMeaningfulData(sm)
	}
	t.Log(len(codes))
}
