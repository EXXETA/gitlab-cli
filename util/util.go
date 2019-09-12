package util

import (
	"encoding/json"
	"fmt"
)

// PrintPrettyJSON prints json object in pretty form
func PrintPrettyJSON(object interface{}) {
	if object != nil {
		jsonObj, err := json.MarshalIndent(object, "", "  ")
		if err != nil {
			// cannot conver to pretty json
			str := fmt.Sprintf("%#v", object)
			fmt.Println(str)
		}
		fmt.Println(string(jsonObj))
	}
}
