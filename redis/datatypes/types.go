package datatypes

import (
	"encoding/json"
	"errors"
	"log"
	"redis/enums"
	"strconv"
	"strings"
)

type ObjectSetType struct {
	Key      string      `json:"key"`
	Val      interface{} `json:"val"`
	DataType string      `json:"type"`
}

func (valObj *ObjectSetType) toString() (string, error) {
	if str, ok := valObj.Val.(string); ok {
		return str, nil
	}
	strBytes, err := json.Marshal(valObj.Val)
	return string(strBytes), err
}

func GetRedisTypeObject(val, type_ string) (RedisType, error) {
	switch enums.GetTypeEnum(type_) {
	case enums.INT:
		valInt, err := strconv.ParseInt(val, 0, 64)
		if err != nil {
			return nil, err
		}
		return &IntType{
			BaseType: BaseType[int64]{
				Type: type_,
				Val:  valInt,
			},
		}, nil
	case enums.FLOAT:
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		return &FloatType{
			BaseType: BaseType[float64]{
				Type: type_,
				Val:  valFloat,
			},
		}, nil
	case enums.STRING:
		return &StringType{
			BaseType: BaseType[string]{
				Type: type_,
				Val:  val,
			},
		}, nil
	case enums.LIST:
		list_val := make([]RedisType, 0)
		var values []ObjectSetType
		err := json.Unmarshal([]byte(val), &values)
		if err != nil {
			return nil, err
		}
		for _, v := range values {
			val, err := v.toString()
			if err != nil {
				return nil, err
			}
			newVal, err := GetRedisTypeObject(val, v.DataType)
			if err != nil {
				return nil, err
			}
			list_val = append(list_val, newVal)
		}
		return &ListType{
			BaseType: BaseType[[]RedisType]{
				Type: type_,
				Val:  list_val,
			},
		}, nil
	case enums.SET:
		set_map := make(map[string]bool)
		for _, v := range strings.Split(val, ",") {
			set_map[v] = true
		}
		return &SetType{
			BaseType: BaseType[map[string]bool]{
				Type: type_,
				Val:  set_map,
			},
		}, nil
	case enums.OBJECT:
		set_map := make(map[string]RedisType)
		var values []ObjectSetType
		err := json.Unmarshal([]byte(val), &values)
		if err != nil {
			log.Println("error in json unmarshall:", val)
			return nil, err
		}
		for _, v := range values {
			val, err := v.toString()
			if err != nil {
				return nil, err
			}
			log.Println("key:", v.Key, "val:", val)
			newVal, err := GetRedisTypeObject(val, v.DataType)
			if err != nil {
				return nil, err
			}
			set_map[v.Key] = newVal
		}
		return &ObjectType{
			BaseType: BaseType[map[string]RedisType]{
				Type: type_,
				Val:  set_map,
			},
		}, nil
	default:
		return nil, errors.New("type does not matched")
	}

}
