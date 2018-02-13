package igmedia

import (
	"testing"
)

func TestStripQueryString(t *testing.T) {
	u := stripQueryString("https://example.com/myvideo.mp4?abc=d")
	if u != "https://example.com/myvideo.mp4" {
		t.Error(u)
	}
}
