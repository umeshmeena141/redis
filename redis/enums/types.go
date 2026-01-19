package enums

type Types int

const (
	INT Types = iota
	FLOAT
	STRING
	LIST
	SET
	OBJECT
)

var ENUM_MAPPING = map[string]Types{
	"int":    INT,
	"float":  FLOAT,
	"string": STRING,
	"list":   LIST,
	"set":    SET,
	"dict":   OBJECT,
}

func GetTypeEnum(type_ string) Types {
	typeEnum, exists := ENUM_MAPPING[type_]
	if exists {
		return typeEnum
	}
	return STRING
}
