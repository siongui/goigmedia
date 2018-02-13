package igmedia

import (
	"encoding/json"
	"fmt"
	"strings"
)

const urlPost = `https://www.instagram.com/p/{{CODE}}/?__a=1`

type postInfo struct {
	GraphQL struct {
		ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

type ShortcodeMedia struct {
	EdgeMedia
	EdgeSidecarToChildren struct {
		Edges []struct {
			Nodes EdgeMedia `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`
}

type EdgeMedia struct {
	Typename         string `json:"__typename"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`
	VideoUrl         string `json:"video_url"`
	TakenAtTimestamp int64  `json:"taken_at_timestamp"`
}

// return URL of image with best resolution
func (em *EdgeMedia) getImageUrl() string {
	res := em.DisplayResources
	return res[len(res)-1].Src
}

func (em *EdgeMedia) getVideoUrl() string {
	return em.VideoUrl
}

func printMeaningfulData(sm ShortcodeMedia) {
	switch sm.Typename {
	case "GraphImage":
		fmt.Println(sm.getImageUrl())
	case "GraphVideo":
		fmt.Println(sm.getVideoUrl())
	case "GraphSidecar":
		fmt.Println("")
	default:
		panic(sm.Typename)
	}
	printTimestamp(sm.TakenAtTimestamp)
}

// Given the code of the post, return url of the post.
func codeToUrl(code string) string {
	return strings.Replace(urlPost, "{{CODE}}", code, 1)
}

// Given code of post, return information of the post with login status.
func (m *IGApiManager) GetPostInfo(code string) (sm ShortcodeMedia, err error) {
	url := codeToUrl(code)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)

	pi := postInfo{}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		return
	}
	sm = pi.GraphQL.ShortcodeMedia
	return
}
