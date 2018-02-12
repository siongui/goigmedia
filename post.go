package igmedia

import (
	"encoding/json"
	"fmt"
	"strings"
)

const urlPost = `https://www.instagram.com/p/{{CODE}}/?__a=1`

type PostInfo struct {
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

func getBestResolutionImageUrl(pi PostInfo) string {
	res := pi.GraphQL.ShortcodeMedia.DisplayResources
	return res[len(res)-1].Src
}

func getVideoUrl(pi PostInfo) string {
	return pi.GraphQL.ShortcodeMedia.VideoUrl
}

func printMeaningfulData(pi PostInfo) {
	switch pi.GraphQL.ShortcodeMedia.Typename {
	case "GraphImage":
		fmt.Println(getBestResolutionImageUrl(pi))
	case "GraphVideo":
		fmt.Println(getVideoUrl(pi))
	case "GraphSidecar":
		fmt.Println("")
	default:
		panic(pi.GraphQL.ShortcodeMedia.Typename)
	}
	printTimestamp(pi.GraphQL.ShortcodeMedia.TakenAtTimestamp)
}

// Given the code of the post, return url of the post.
func codeToUrl(code string) string {
	return strings.Replace(urlPost, "{{CODE}}", code, 1)
}

// Given code of post, return information of the post with login status.
func (m *IGApiManager) GetPostInfo(code string) (pi PostInfo, err error) {
	url := codeToUrl(code)
	fmt.Println(url) // to be deleted
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)

	err = json.Unmarshal(b, &pi)
	return
}
