package datatypes

import (
	"encoding/json"
	"errors"
	"log"
	"redis/enums"
)

type RedisType interface {
	Get(key string) string
	Set(key, value, Type string) string
	Del(key string) string
	Json() string
	FindByPath(key string) (RedisType, error)
	GetType() string
}

type BaseType[T any] struct {
	Type string
	Val  T
}

func (valObj *BaseType[T]) Get(key string) string {
	log.Println("Can not get values for " + valObj.Type + " type")
	return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
}

func (valObj *BaseType[T]) Set(key, value, Type string) string {
	log.Println("Can not set values for " + valObj.Type + " type")
	return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
}

func (valObj *BaseType[T]) Del(key string) string {
	log.Println("Can not delete values for " + valObj.Type + " type")
	return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
}

func (valObj *BaseType[T]) FindByPath(key string) (RedisType, error) {
	log.Println("Can not get values for " + valObj.Type + " type")
	return nil, errors.New("key doesn't exists")
}

func (valObj *BaseType[T]) Json() string {
	val, err := json.Marshal(valObj.Val)
	if err != nil {
		log.Println("Error while converting "+valObj.Type+" value to json", valObj.Val)
		return enums.GetResponseMessage(enums.ERROR)
	}
	return string(val)
}

func (valObj *BaseType[T]) GetType() string {
	return valObj.Type
}
