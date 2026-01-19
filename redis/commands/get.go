package commands

import (
	"log"
	d "redis/datatypes"
)

type GetCommandStruct struct {
	BASE_CMD
}

func (cmd *GetCommandStruct) HandleCommand() string {
	log.Printf("Get command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		return objToUpdate.Get(key)
	}
	return cmd.executeCommand(action)
}
