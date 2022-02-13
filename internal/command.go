package fzg

import (
	"fmt"
	"reflect"
)

type Command interface{}

func parseCommand(commandConf Command) []string {
	if commandConf == nil {
		return []string{}
	}
	command := []string{}
	switch commandConf.(type) {
	case []interface{}:
		return commandFromSequence(commandConf)
	case string:
		return commandFromString(commandConf)
	default:
		return command
	}
}

func commandFromSequence(ivalue interface{}) []string {
	strSlice := []string{}
	v := reflect.ValueOf(ivalue)
	for i := 0; i < v.Len(); i++ {
		switch ivalue := v.Index(i).Interface().(type) {
		case []interface{}:
			for _, element := range ivalue {
				strSlice = append(strSlice, element.(string))
			}
		case string:
			strSlice = append(strSlice, fmt.Sprint(v.Index(i)))
		}
	}
	return strSlice
}

func commandFromString(ivalue interface{}) []string {
	return []string{fmt.Sprint(ivalue)}
}
