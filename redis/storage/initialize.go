package storage

import (
	"net"
	d "redis/datatypes"
)

func InitializeStore() *Storage {
	store := Storage{
		Tables: map[string]*Table{
			"default": &Table{
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
			},
		},
	}
	go PublishChannelMessage(store.Tables["default"], "updates")
	return &store
}
