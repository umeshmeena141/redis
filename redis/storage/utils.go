package storage

import "fmt"

func PublishChannelMessage(table *Table, chanName string) {
	for {
		msg := <-table.Channels[chanName]
		for _, conn := range table.ChannelsConnMap[chanName] {
			fmt.Fprintln(conn, msg)
		}
	}
}
