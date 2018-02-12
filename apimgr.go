// Package igmedia helps you get all URLs of posts of a specific Instagram user,
// and also media (photos and videos) links of posts.
package igmedia

type IGApiManager struct {
	dsUserId  string
	sessionid string
	csrftoken string
}

// After login to Instagram, you can get the cookies of *ds_user_id*,
// *sessionid*, *csrftoken* in Chrome Developer Tools.
// See https://stackoverflow.com/a/44773079
// or
// https://github.com/hoschiCZ/instastories-backup#obtain-cookies
func NewInstagramApiManager(ds_user_id, sessionid, csrftoken string) *IGApiManager {
	return &IGApiManager{
		dsUserId:  ds_user_id,
		sessionid: sessionid,
		csrftoken: csrftoken,
	}
}
