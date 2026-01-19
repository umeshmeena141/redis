package datatypes

import "encoding/json"

type SetType struct {
	BaseType[map[string]bool]
}

func (valObj *SetType) Json() string {
	jsonObj := make([]string, 0)
	for key, _ := range valObj.Val {
		jsonObj = append(jsonObj, key)
	}
	val, _ := json.Marshal(jsonObj)
	return string(val)
}
