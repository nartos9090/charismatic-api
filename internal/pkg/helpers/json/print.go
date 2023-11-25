package helpers_json

import (
	"encoding/json"
	"fmt"
)

func Print(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
