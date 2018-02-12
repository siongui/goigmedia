package igmedia

import (
	"encoding/json"
	"fmt"
	"time"
)

func JsonPrettyPrint(v interface{}) {
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
