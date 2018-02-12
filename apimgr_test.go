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

	ui, err := mgr.GetUserInfo(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, node := range ui.Media.Nodes {
		sm, err := mgr.GetPostInfo(node.Code)
		if err != nil {
			t.Error(err)
			return
		}
		//JsonPrettyPrint(sm)
		printMeaningfulData(sm)
	}
}
