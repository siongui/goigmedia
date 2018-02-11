// Package goiguserid returns id of Instagram user given the user name.
package igmedia

// Get Instagram user information, such as id and biography, via username.
// See ``Instagram API -Get the userId - Stack Overflow``
// https://stackoverflow.com/a/44773079

import (
	"encoding/json"
	"net/http"
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

// Given user name, return information of the user name.
// Currently only id and biography is returned.
func GetUserInfo(username string) (ui UserInfo, err error) {
	url := strings.Replace(UrlUserInfo, "{{USERNAME}}", username, 1)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	r := RawUserResp{}
	if err = dec.Decode(&r); err != nil {
		return
	}
	ui = r.User
	return
}

// Given user name, return id of the user name.
func GetUserId(username string) (id string, err error) {
	ui, err := GetUserInfo(username)
	if err != nil {
		return
	}
	id = ui.Id
	return
}
