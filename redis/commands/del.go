package commands

import (
	"log"
	d "redis/datatypes"
)

type DelCommandStruct struct {
	BASE_CMD
}

func (cmd *DelCommandStruct) HandleCommand() string {
	log.Printf("Del command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		return objToUpdate.Del(key)
	}
	return cmd.executeCommand(action)
}
