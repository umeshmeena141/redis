package commands

import (
	"encoding/json"
	"log"
	d "redis/datatypes"
	"redis/enums"
	"strconv"
)

func performListAction(dictObj d.RedisType, key string, action func(*d.ListType) string) string {
	obj, ok := dictObj.(*d.ObjectType)
	if !ok {
		return enums.GetResponseMessage(enums.ERROR)
	}
	objToUpdate, exists := obj.Val[key]
	if !exists {
		return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
	}
	if objToUpdate.GetType() == "list" {
		listObj, ok := objToUpdate.(*d.ListType)
		if ok {
			return action(listObj)
		}
		return enums.GetResponseMessage(enums.ERROR)
	}
	return enums.GetResponseMessage(enums.ERROR)
}

type LPUSHCommandStruct struct {
	BASE_CMD
}

func (cmd *LPUSHCommandStruct) HandleCommand() string {
	log.Printf("LPUSH command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		listAction := func(listObj *d.ListType) string {
			var pushVal map[string]string
			err := json.Unmarshal([]byte(cmd.Args[2]), &pushVal)
			if err != nil {
				return enums.GetResponseMessage(enums.ERROR)
			}
			redisTypeObj, err := d.GetRedisTypeObject(pushVal["val"], pushVal["type"])
			if err != nil {
				return enums.GetResponseMessage(enums.ERROR)
			}
			listObj.Val = append(listObj.Val, redisTypeObj)
			return enums.GetResponseMessage(enums.OK)
		}
		return performListAction(objToUpdate, key, listAction)
	}
	return cmd.executeCommand(action)
}

type LPOPCommandStruct struct {
	BASE_CMD
}

func (cmd *LPOPCommandStruct) HandleCommand() string {
	log.Printf("Pop from list command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		listAction := func(listObj *d.ListType) string {
			popObj := listObj.Val[len(listObj.Val)-1]
			listObj.Val = listObj.Val[:len(listObj.Val)-1]
			return popObj.Json()
		}
		return performListAction(objToUpdate, key, listAction)
	}
	return cmd.executeCommand(action)
}

type LGETCommandStruct struct {
	BASE_CMD
}

func (cmd *LGETCommandStruct) HandleCommand() string {
	log.Printf("Get from list command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		listAction := func(listObj *d.ListType) string {
			index, _ := strconv.ParseInt(cmd.Args[2], 0, 64)
			return listObj.Val[index].Json()
		}
		return performListAction(objToUpdate, key, listAction)
	}
	return cmd.executeCommand(action)
}
