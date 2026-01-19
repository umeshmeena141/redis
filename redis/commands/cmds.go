package commands

import (
	"fmt"
	"log"
	"net"
	d "redis/datatypes"
	"redis/enums"
	store "redis/storage"
	"slices"
	"strings"
)

type CMD interface {
	HandleCommand() string
	checkArgs() bool
	checkForPath() (d.RedisType, string, string)
	getTable() (*store.Table, string)
	executeCommand(actionFunc func(d.RedisType, string) string) string
}

type BASE_CMD struct {
	Name            string
	ExpectedArgsLen int
	Args            []string
	DataStore       *store.Storage
	TableName       *string
	Connection      net.Conn
}

func (cmd *BASE_CMD) checkArgs() bool {
	if len(cmd.Args) != cmd.ExpectedArgsLen {
		log.Println("Invalid args for: ", cmd.Name, cmd.ExpectedArgsLen, len(cmd.Args))
		return false
	}
	return true
}

func (cmd *BASE_CMD) HandleCommand() string {
	log.Printf("Base command recieved")
	return enums.GetResponseMessage(enums.UNKNOWN_COMMAND)
}

func (cmd *BASE_CMD) getTable() (*store.Table, string) {
	table, exists := cmd.DataStore.Tables[*cmd.TableName]
	if !exists {
		return nil, enums.GetResponseMessage(enums.TABLE_NOT_EXISTS)
	}
	return table, ""
}

func (cmd *BASE_CMD) checkForPath() (d.RedisType, string, string) {
	table, _ := cmd.getTable()
	objToUpdate, err := table.Data.FindByPath(cmd.Args[1])
	if err != nil {
		return nil, "", enums.GetResponseMessage(enums.INVALID_PATH_ERR)
	}
	keys := strings.Split(cmd.Args[1], ".")
	key := keys[len(keys)-1]
	return objToUpdate, key, ""
}

func sendUpdates(table *store.Table, commandName, path string) {
	commandsToNotSendUpdates := []string{"GET", "LGET", "SFIND", "PING"}
	if !slices.Contains(commandsToNotSendUpdates, commandName) {
		msg := fmt.Sprintf("{\"CMD\": \"%s\", \"Path\":\"%s\"}", commandName, path)
		table.Channels["updates"] <- msg
	}
}

func (cmd *BASE_CMD) executeCommand(actionFunc func(d.RedisType, string) string) string {
	if cmd.checkArgs() {
		table, err := cmd.getTable()
		if err != "" {
			return err
		}
		objToUpdate, key, errorMsg := cmd.checkForPath()
		if errorMsg != "" {
			return errorMsg
		}
		table.Mutex.Lock()
		defer sendUpdates(table, cmd.Name, cmd.Args[1])
		defer table.Mutex.Unlock()
		return actionFunc(objToUpdate, key)
	}
	return enums.GetResponseMessage(enums.INVALID_ARGS)
}

func GetCMD(conn net.Conn, dataStore *store.Storage, tableName *string, args []string) CMD {
	baseType := BASE_CMD{
		Name:       args[0],
		Args:       args,
		DataStore:  dataStore,
		TableName:  tableName,
		Connection: conn,
	}
	switch enums.GetCommand(args[0]) {
	case enums.GET:
		baseType.ExpectedArgsLen = 2
		return &GetCommandStruct{BASE_CMD: baseType}
	case enums.SET_COMMAND:
		baseType.ExpectedArgsLen = 4
		return &SetCommandStruct{BASE_CMD: baseType}
	case enums.DEL:
		baseType.ExpectedArgsLen = 2
		return &DelCommandStruct{BASE_CMD: baseType}
	case enums.PING:
		baseType.ExpectedArgsLen = 1
		return &PingCommandStruct{BASE_CMD: baseType}
	case enums.SET_TABLE:
		baseType.ExpectedArgsLen = 2
		return &SetTableCommandStruct{BASE_CMD: baseType}
	case enums.CREATE_TABLE:
		baseType.ExpectedArgsLen = 2
		return &CreateTableCommandStruct{BASE_CMD: baseType}
	case enums.INC:
		baseType.ExpectedArgsLen = 2
		return &INCCommandStruct{BASE_CMD: baseType}
	case enums.INC_N:
		baseType.ExpectedArgsLen = 3
		return &INCNCommandStruct{BASE_CMD: baseType}
	case enums.DEC:
		baseType.ExpectedArgsLen = 2
		return &DECCommandStruct{BASE_CMD: baseType}
	case enums.DEC_N:
		baseType.ExpectedArgsLen = 3
		return &DECNCommandStruct{BASE_CMD: baseType}

	case enums.LPUSH:
		baseType.ExpectedArgsLen = 3
		return &LPUSHCommandStruct{BASE_CMD: baseType}
	case enums.LPOP:
		baseType.ExpectedArgsLen = 2
		return &LPOPCommandStruct{BASE_CMD: baseType}
	case enums.LGET:
		baseType.ExpectedArgsLen = 3
		return &LGETCommandStruct{BASE_CMD: baseType}

	case enums.SADD:
		baseType.ExpectedArgsLen = 3
		return &SADDCommandStruct{BASE_CMD: baseType}
	case enums.SPOP:
		baseType.ExpectedArgsLen = 3
		return &SPOPCommandStruct{BASE_CMD: baseType}
	case enums.SFIND:
		baseType.ExpectedArgsLen = 3
		return &SFINDCommandStruct{BASE_CMD: baseType}
	case enums.CREATE_CHAN:
		baseType.ExpectedArgsLen = 2
		return &CreateChannelCommandStruct{BASE_CMD: baseType}
	case enums.SUB_CHAN:
		baseType.ExpectedArgsLen = 2
		return &SubscribeChannelCommandStruct{BASE_CMD: baseType}
	case enums.PUB_CHAN:
		baseType.ExpectedArgsLen = 3
		return &PublishChannelCommandStruct{BASE_CMD: baseType}

	default:
		baseType.ExpectedArgsLen = 0
		return &baseType
	}
}
