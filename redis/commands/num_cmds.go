package commands

import (
	"log"
	d "redis/datatypes"
	"redis/enums"
	"strconv"
)

func increaseOrDecreaseAction(dictObj d.RedisType, key string, n int64) string {
	obj, ok := dictObj.(*d.ObjectType)
	if !ok {
		return enums.GetResponseMessage(enums.ERROR)
	}
	objToUpdate, exists := obj.Val[key]
	if !exists {
		return enums.GetResponseMessage(enums.INVALID_PATH_ERR)
	}
	log.Println("Type of object: ", objToUpdate.GetType())
	if objToUpdate.GetType() == "int" {
		intObj, ok := objToUpdate.(*d.IntType)
		if ok {
			intObj.Val += n
			return enums.GetResponseMessage(enums.OK)
		}
		return enums.GetResponseMessage(enums.ERROR)
	}
	return enums.GetResponseMessage(enums.ERROR)
}

type INCCommandStruct struct {
	BASE_CMD
}

func (cmd *INCCommandStruct) HandleCommand() string {
	log.Printf("Increase command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		return increaseOrDecreaseAction(objToUpdate, key, 1)
	}
	return cmd.executeCommand(action)
}

type INCNCommandStruct struct {
	BASE_CMD
}

func (cmd *INCNCommandStruct) HandleCommand() string {
	log.Printf("Increase By N command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		inc, _ := strconv.ParseInt(cmd.Args[2], 0, 64)
		return increaseOrDecreaseAction(objToUpdate, key, inc)
	}
	return cmd.executeCommand(action)
}

type DECCommandStruct struct {
	BASE_CMD
}

func (cmd *DECCommandStruct) HandleCommand() string {
	log.Printf("Decrease command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		return increaseOrDecreaseAction(objToUpdate, key, -1)
	}
	return cmd.executeCommand(action)
}

type DECNCommandStruct struct {
	BASE_CMD
}

func (cmd *DECNCommandStruct) HandleCommand() string {
	log.Printf("Decrease By N command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		inc, _ := strconv.ParseInt(cmd.Args[2], 0, 64)
		return increaseOrDecreaseAction(objToUpdate, key, -inc)
	}
	return cmd.executeCommand(action)
}
