// Package goiguserid returns id of Instagram user given the user name.
package igmedia

// Get Instagram user information, such as id and biography, via username.
// See ``Instagram API -Get the userId - Stack Overflow``
// https://stackoverflow.com/a/44773079

import (
	"encoding/json"
	"strings"
)

// no need to login or cookie to access this URL
const UrlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

// used to decode the JSON data
type RawUserResp struct {
	User UserInfo
}

// You can add more fields in the struct to get more information
// See response/types.go in github.com/ahmdrz/goinsta
type UserInfo struct {
	Id        string `json:"id"`
	Biography string `json:"biography"`
}

// Given user name, return information of the user name without login.
// Currently only id and biography is returned.
func GetUserInfoNoLogin(username string) (ui UserInfo, err error) {
	url := strings.Replace(UrlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponseNoLogin(url)

	r := RawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.User
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
