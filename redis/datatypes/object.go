package datatypes

import (
	"encoding/json"
	"errors"
	"log"
	"redis/enums"
	"strings"
)

type ObjectType struct {
	BaseType[map[string]RedisType]
}

func (valObj *ObjectType) Get(key string) string {
	if valObj.Val[key] == nil {
		return "null"
	}
	return valObj.Val[key].Json()
}

func (valObj *ObjectType) Set(key, value, type_ string) string {
	typedVal, err := GetRedisTypeObject(value, type_)
	if err != nil {
		return enums.GetResponseMessage(enums.INVALID_TYPE_ERR)
	}
	valObj.Val[key] = typedVal
	return enums.GetResponseMessage(enums.OK)
}

func (valObj *ObjectType) Del(key string) string {
	delete(valObj.Val, key)
	return enums.GetResponseMessage(enums.OK)
}

func (valObj *ObjectType) FindByPath(key string) (RedisType, error) {
	if key == "" {
		log.Println("Not a valid key")
		return nil, errors.New("key doesn't exists")
	} else if strings.Contains(key, ".") {
		splits := strings.Split(key, ".")
		newKey := strings.Join(splits[1:], ".")
		newValObj, exists := valObj.Val[splits[0]]
		if !exists {
			return nil, errors.New("key doesn't exists")
		}
		return newValObj.FindByPath(newKey)
	} else if strings.Contains(key, "[") && strings.Contains(key, "]") {
		log.Println("Not a valid key")
		return nil, errors.New("key doesn't exists")
	}
	log.Println("Found valid object for the path")
	return valObj, nil
}

func (valObj *ObjectType) Json() string {
	jsonObj := make(map[string]string)
	for key, val := range valObj.Val {
		jsonObj[key] = val.Json()
	}
	val, _ := json.Marshal(jsonObj)
	return string(val)
}
