package fzg

import (
	"fmt"
	"reflect"
	"strings"
)

type Command interface{}

func parseCommand(commandConf Command) string {
	if commandConf == nil {
		return ""
	}
	command := ""
	switch commandConf.(type) {
	case []interface{}:
		return commandFromSequence(commandConf)
	case string:
		return commandFromString(commandConf)
	default:
		return command
	}
}

func commandFromSequence(ivalue interface{}) string {
	strSlice := []string{}
	v := reflect.ValueOf(ivalue)
	for i := 0; i < v.Len(); i++ {
		strSlice = append(strSlice, fmt.Sprint(v.Index(i)))
	}
	return strings.Join(strSlice, " ")
}

func commandFromString(ivalue interface{}) string {
	return fmt.Sprint(ivalue)
}
