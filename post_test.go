package igmedia

import (
	"os"
	"testing"
)

func TestGetPostInfo(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	sm, err := mgr.GetPostInfo(os.Getenv("IG_TEST_CODE"))
	if err != nil {
		t.Error(err)
		return
	}
	jsonPrettyPrint(sm)
	printMeaningfulData(sm)
}
