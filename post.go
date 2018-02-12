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
	TakenAtTimestamp int64 `json:"taken_at_timestamp"`
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
