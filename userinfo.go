// Package goiguserid returns id of Instagram user given the user name.
package igmedia

// Get Instagram user information, such as id and biography, via username.
// See ``Instagram API -Get the userId - Stack Overflow``
// https://stackoverflow.com/a/44773079

import (
	"encoding/json"
	"strings"
)

// no need to login or cookie to access this URL. But if login to Instagram,
// private account will return private data if you are allowed to view the
// private account.
const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

// used to decode the JSON data
type RawUserResp struct {
	User UserInfo `json:"user"`
}

// You can add more fields in the struct to get more information
// See response/types.go in github.com/ahmdrz/goinsta
type UserInfo struct {
	Biography       string `json:"biography"`
	ExternalUrl     string `json:"external_url"`
	FullName        string `json:"full_name"`
	Id              string `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	ProfilePicUrlHd string `json:"profile_pic_url_hd"`
	Username        string `json:"username"`
	Media           struct {
		Nodes    []MediaNode `json:"nodes"`
		Count    int64       `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
	} `json:"media"`
}

type MediaNode struct {
	Code    string `json:"code"` // url of the post
	Date    int64  `json:"date"`
	Caption string `json:"caption"`
}

// Given user name, return information of the user name without login.
func GetUserInfoNoLogin(username string) (ui UserInfo, err error) {
	url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	r := RawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.User
	return
}

// Given user name, return information of the user name.
func (m *IGApiManager) GetUserInfo(username string) (ui UserInfo, err error) {
	url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	r := RawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.User
	return
}

// Given user name, return codes of all posts of the user.
func (m *IGApiManager) GetAllPostCode(username string) (codes []string, err error) {
	ui, err := m.GetUserInfo(username)
	if err != nil {
		return
	}
	for _, node := range ui.Media.Nodes {
		codes = append(codes, node.Code)
	}
	for ui.Media.PageInfo.HasNextPage == true {
		url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
		url = url + "&max_id=" + ui.Media.PageInfo.EndCursor
		b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
		if err != nil {
			return codes, err
		}
		r := RawUserResp{}
		if err = json.Unmarshal(b, &r); err != nil {
			return codes, err
		}
		ui = r.User
		for _, node := range ui.Media.Nodes {
			codes = append(codes, node.Code)
		}
	}
	return
}

// Given user name, return id of the user name.
func GetUserId(username string) (id string, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}
	id = ui.Id
	return
}
