package igmedia

import (
	"encoding/json"
	"strings"
)

// no need to login or cookie to access this URL. But if login to Instagram,
// private account will return private data if you are allowed to view the
// private account.
const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

// used to decode the JSON data
type rawUserResp struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Biography       string `json:"biography"`
	ExternalUrl     string `json:"external_url"`
	FullName        string `json:"full_name"`
	Id              string `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	ProfilePicUrlHd string `json:"profile_pic_url_hd"`
	Username        string `json:"username"`
	Media           struct {
		Nodes []struct {
			Code    string `json:"code"` // url of the post
			Date    int64  `json:"date"`
			Caption string `json:"caption"`
		} `json:"nodes"`
		Count    int64 `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
	} `json:"media"`
}

// Given user name, return information of the user name without login.
func GetUserInfoNoLogin(username string) (ui UserInfo, err error) {
	url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	r := rawUserResp{}
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

	r := rawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.User
	return
}

// Given user name, return codes of all posts of the user.
// TODO: add sleep at the end of forloop. If the number of posts is over 2400,
// Instagram API will return http response code 429 (Too Many Requests)
func (m *IGApiManager) GetAllPostCode(username string) (codes []string, err error) {
	r := rawUserResp{}
	r.User.Media.PageInfo.HasNextPage = true
	for r.User.Media.PageInfo.HasNextPage == true {
		url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
		if len(codes) != 0 {
			url = url + "&max_id=" + r.User.Media.PageInfo.EndCursor
		}

		b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
		if err != nil {
			return codes, err
		}

		if err = json.Unmarshal(b, &r); err != nil {
			return codes, err
		}

		for _, node := range r.User.Media.Nodes {
			codes = append(codes, node.Code)
		}
		printPostCount(len(codes), url)
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
