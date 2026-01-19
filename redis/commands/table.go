package commands

import (
	"log"
	"net"
	d "redis/datatypes"
	"redis/enums"
	store "redis/storage"
)

type SetTableCommandStruct struct {
	BASE_CMD
}

func (cmd *SetTableCommandStruct) HandleCommand() string {
	log.Printf("Set Table command recieved")
	if cmd.checkArgs() {
		_, exists := cmd.DataStore.Tables[cmd.Args[1]]
		if !exists {
			return enums.GetResponseMessage(enums.TABLE_NOT_EXISTS)
		}
		*cmd.TableName = cmd.Args[1]
		return enums.GetResponseMessage(enums.OK)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}

type CreateTableCommandStruct struct {
	BASE_CMD
}

func (cmd *CreateTableCommandStruct) HandleCommand() string {
	log.Printf("Create Table command recieved")
	if cmd.checkArgs() {
		_, exists := cmd.DataStore.Tables[cmd.Args[1]]
		if exists {
			return enums.GetResponseMessage(enums.TABLE_ALREADY_EXISTS)
		}
		cmd.DataStore.Tables[cmd.Args[1]] = &store.Table{
			Data: d.ObjectType{
				BaseType: d.BaseType[map[string]d.RedisType]{
					Type: "dict",
					Val:  make(map[string]d.RedisType),
				},
			},
			Channels: map[string]chan string{
				"updates": make(chan string, 1),
			},
			ChannelsConnMap: map[string][]net.Conn{
				"updates": make([]net.Conn, 0),
			},
		}
		go store.PublishChannelMessage(cmd.DataStore.Tables[cmd.Args[1]], "updates")
		*cmd.TableName = cmd.Args[1]
		return enums.GetResponseMessage(enums.OK)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}
