package datatypes

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
)

type ListType struct {
	BaseType[[]RedisType]
}

func (valObj *ListType) FindByPath(key string) (RedisType, error) {
	if strings.Contains(key, "[") && strings.Contains(key, "]") {
		splits := strings.Split(key, ".")
		indexStr := strings.TrimSuffix(strings.TrimPrefix(splits[0], "["), "]")
		index, err := strconv.ParseInt(indexStr, 0, 32)
		if err != nil || index < 0 || int(index) >= len(valObj.Val) {
			return nil, errors.New("index doesn't exists")
		}
		newKey := strings.Join(splits[1:], ".")
		return valObj.Val[index].FindByPath(newKey)
	}
	log.Println("Not a valid key")
	return nil, errors.New("key doesn't exists")
}

func (valObj *ListType) Json() string {
	jsonObj := make([]string, 0)
	for _, elem := range valObj.Val {
		val := elem.Json()
		jsonObj = append(jsonObj, val)
	}
	val, _ := json.Marshal(jsonObj)
	return string(val)
}
