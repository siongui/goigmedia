package igmedia

import (
	"encoding/json"
	"fmt"
)

func JsonPrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
