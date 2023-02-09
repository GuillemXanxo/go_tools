package backend

import (
	"encoding/json"
	"fmt"
)

func PrintQuery(src interface{}) {
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
