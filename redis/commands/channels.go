package commands

import (
	"log"
	"net"
	"redis/enums"
	"redis/storage"
)

type CreateChannelCommandStruct struct {
	BASE_CMD
}

func createChannelCoroutine(table *storage.Table, chanName string) {
	channel := make(chan string, 1)
	table.Channels[chanName] = channel
	table.ChannelsConnMap[chanName] = make([]net.Conn, 0)
	storage.PublishChannelMessage(table, chanName)
}

func (cmd *CreateChannelCommandStruct) HandleCommand() string {
	log.Printf("Create channel command recieved")
	if cmd.checkArgs() {
		table, _ := cmd.getTable()
		table.Mutex.Lock()
		defer table.Mutex.Unlock()
		_, exists := table.Channels[cmd.Args[1]]
		if exists {
			return enums.GetResponseMessage(enums.CHANNEL_ALREADY_EXISTS)
		}
		go createChannelCoroutine(table, cmd.Args[1])

		return enums.GetResponseMessage(enums.OK)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}

type SubscribeChannelCommandStruct struct {
	BASE_CMD
}

func (cmd *SubscribeChannelCommandStruct) HandleCommand() string {
	log.Printf("Subscribe channel command recieved")
	if cmd.checkArgs() {
		table, _ := cmd.getTable()
		table.Mutex.Lock()
		defer table.Mutex.Unlock()
		_, exists := table.Channels[cmd.Args[1]]
		if !exists {
			return enums.GetResponseMessage(enums.CHANNEL_NOT_EXISTS)
		}
		table.ChannelsConnMap[cmd.Args[1]] = append(table.ChannelsConnMap[cmd.Args[1]], cmd.Connection)
		return enums.GetResponseMessage(enums.OK)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}

type PublishChannelCommandStruct struct {
	BASE_CMD
}

func (cmd *PublishChannelCommandStruct) HandleCommand() string {
	log.Printf("Publish channel command recieved")
	if cmd.checkArgs() {
		table, _ := cmd.getTable()
		table.Mutex.Lock()
		defer table.Mutex.Unlock()
		_, exists := table.Channels[cmd.Args[1]]
		if !exists {
			return enums.GetResponseMessage(enums.CHANNEL_NOT_EXISTS)
		}
		table.Channels[cmd.Args[1]] <- cmd.Args[2]
		return enums.GetResponseMessage(enums.OK)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}
