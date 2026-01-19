package storage

import (
	"net"
	"redis/datatypes"
	"sync"
)

type Storage struct {
	Tables map[string]*Table
}

type Table struct {
	Mutex           sync.Mutex
	Data            datatypes.ObjectType
	Channels        map[string]chan string
	ChannelsConnMap map[string][]net.Conn
}
