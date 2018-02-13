package igmedia

import (
	"encoding/json"
	"fmt"
	"strings"
)

const urlPost = `https://www.instagram.com/p/{{CODE}}/?__a=1`

type postInfo struct {
	GraphQL struct {
		ShortcodeMedia EdgeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

type EdgeMedia struct {
	Typename   string `json:"__typename"`
	Shortcode  string `json:"shortcode"`
	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`
	DisplayUrl       string `json:"display_url"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`
	VideoUrl         string `json:"video_url"`
	IsVideo          bool   `json:"is_video"`
	TakenAtTimestamp int64  `json:"taken_at_timestamp"`
	Location         struct {
		Id            string `json:"id"`
		HasPublicPage bool   `json:"has_public_page"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
	} `json:"location"`
	EdgeSidecarToChildren struct {
		Edges []struct {
			Nodes EdgeMedia `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`
}

// return URL of image with best resolution
func (em *EdgeMedia) getImageUrl() string {
	res := em.DisplayResources
	return res[len(res)-1].Src
}

func (em *EdgeMedia) getVideoUrl() string {
	return em.VideoUrl
}

func printMeaningfulData(em EdgeMedia) {
	switch em.Typename {
	case "GraphImage":
		fmt.Println(em.getImageUrl())
	case "GraphVideo":
		fmt.Println(em.getVideoUrl())
	case "GraphSidecar":
		fmt.Println("")
	default:
		panic(em.Typename)
	}
	printTimestamp(em.TakenAtTimestamp)
}

// Given the code of the post, return url of the post.
func codeToUrl(code string) string {
	return strings.Replace(urlPost, "{{CODE}}", code, 1)
}

// Given code of post, return information of the post with login status.
func (m *IGApiManager) GetPostInfo(code string) (em EdgeMedia, err error) {
	url := codeToUrl(code)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	pi := postInfo{}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		return
	}
	em = pi.GraphQL.ShortcodeMedia
	return
}
