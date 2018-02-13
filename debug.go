package igmedia

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func jsonPrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func printTimestamp(timestamp int64) {
	fmt.Println(formatTimestamp(timestamp))
}

func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func printPostCount(c int, url string) {
	url = strings.Replace(url, "__a=1&", "", 1)
	fmt.Printf("Getting %d from %s ...\n", c, url)
}

// Remove query string in the URL
func stripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}
