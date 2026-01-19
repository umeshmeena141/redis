package commands

import (
	"log"
	d "redis/datatypes"
)

type SetCommandStruct struct {
	BASE_CMD
}

func (cmd *SetCommandStruct) HandleCommand() string {
	log.Printf("Set command recieved")
	action := func(objToUpdate d.RedisType, key string) string {
		return objToUpdate.Set(key, cmd.Args[2], cmd.Args[3])
	}
	return cmd.executeCommand(action)
}
