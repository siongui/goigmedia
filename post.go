package igmedia

import (
	"encoding/json"
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
	TakenAtTimestamp int64 `json:"taken_at_timestamp"`
}

func getBestResolutionUrl(pi PostInfo) string {
	res := pi.GraphQL.ShortcodeMedia.DisplayResources
	return res[len(res)-1].Src
}

// Given the code of the post, return url of the post.
func codeToUrl(code string) string {
	return strings.Replace(urlPost, "{{CODE}}", code, 1)
}

// Given code of post, return information of the post with login status.
func (m *IGApiManager) GetPostInfo(code string) (pi PostInfo, err error) {
	url := codeToUrl(code)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)

	err = json.Unmarshal(b, &pi)
	return
}
