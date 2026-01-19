package commands

import (
	"log"
	d "redis/datatypes"
	"redis/enums"
)

func performSetAction(dictObj d.RedisType, key string, action func(*d.SetType) string) string {
	obj, ok := dictObj.(*d.ObjectType)
	if !ok {
		return enums.GetResponseMessage(enums.ERROR)
	}
	objToUpdate, exists := obj.Val[key]
	if !exists {
		return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
	}
	if objToUpdate.GetType() == "set" {
		setObj, ok := objToUpdate.(*d.SetType)
		if ok {
			return action(setObj)
		}
		return enums.GetResponseMessage(enums.ERROR)
	}
	return enums.GetResponseMessage(enums.ERROR)
}

type SADDCommandStruct struct {
	BASE_CMD
}

func (cmd *SADDCommandStruct) HandleCommand() string {
	log.Printf("SADD command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		setAction := func(setObj *d.SetType) string {
			setObj.Val[cmd.Args[2]] = true
			return enums.GetResponseMessage(enums.OK)
		}
		return performSetAction(objToUpdate, key, setAction)
	}
	return cmd.executeCommand(action)
}

type SPOPCommandStruct struct {
	BASE_CMD
}

func (cmd *SPOPCommandStruct) HandleCommand() string {
	log.Printf("Pop from set command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		setAction := func(setObj *d.SetType) string {
			delete(setObj.Val, cmd.Args[2])
			return enums.GetResponseMessage(enums.OK)
		}
		return performSetAction(objToUpdate, key, setAction)
	}
	return cmd.executeCommand(action)
}

type SFINDCommandStruct struct {
	BASE_CMD
}

func (cmd *SFINDCommandStruct) HandleCommand() string {
	log.Printf("Find element in set command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		setAction := func(setObj *d.SetType) string {
			_, exists := setObj.Val[cmd.Args[2]]
			if exists {
				return "TRUE"
			} else {
				return "FALSE"
			}
		}
		return performSetAction(objToUpdate, key, setAction)
	}
	return cmd.executeCommand(action)
}
