package enums

type Command int

const (
	GET Command = iota
	SET_COMMAND
	DEL
	PING
	SET_TABLE
	CREATE_TABLE
	LPUSH
	LPOP
	LGET
	SADD
	SPOP
	SFIND
	INC
	INC_N
	DEC
	DEC_N
	CREATE_CHAN
	SUB_CHAN
	PUB_CHAN
	UNKNOWN_CMD
)

var COMMAND_MAP = map[string]Command{
	"GET":          GET,
	"SET":          SET_COMMAND,
	"DEL":          DEL,
	"PING":         PING,
	"SET_TABLE":    SET_TABLE,
	"CREATE_TABLE": CREATE_TABLE,
	"UNKNOWN_CMD":  UNKNOWN_CMD,
	"LPUSH":        LPUSH,
	"LPOP":         LPOP,
	"LGET":         LGET,
	"SADD":         SADD,
	"SPOP":         SPOP,
	"SFIND":        SFIND,
	"INC":          INC,
	"INC_N":        INC_N,
	"DEC":          DEC,
	"DEC_N":        DEC_N,
	"CREATE_CHAN":  CREATE_CHAN,
	"SUB_CHAN":     SUB_CHAN,
	"PUB_CHAN":     PUB_CHAN,
}

func GetCommand(command string) Command {
	cmd, exists := COMMAND_MAP[command]
	if exists {
		return cmd
	}
	return UNKNOWN_CMD
}
