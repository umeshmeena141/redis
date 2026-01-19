package commands

import (
	"log"
	"redis/enums"
)

type PingCommandStruct struct {
	BASE_CMD
}

func (cmd *PingCommandStruct) HandleCommand() string {
	log.Printf("Ping command recieved")
	if cmd.checkArgs() {
		return enums.GetResponseMessage(enums.PONG)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}
