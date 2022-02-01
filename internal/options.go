package fzg

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Options map[string]interface{}

func parseOptions(optionsConf Options) string {
	if optionsConf == nil {
		return ""
	}
	options := []string{}
	for name, ivalue := range optionsConf {
		switch ivalue.(type) {
		case map[interface{}]interface{}:
			options = append(options, optionFromMap(name, ivalue))
		case []interface{}:
			options = append(options, optionFromSlice(name, ivalue))
		case bool:
			options = append(options, optionFromBool(name, ivalue))
		case int, float64:
			options = append(options, optionFromNumber(name, ivalue))
		case string:
			options = append(options, optionFromString(name, ivalue))
		}
	}
	sort.Strings(options)
	return strings.Join(options, " ")
}

func optionFromMap(name string, ivalue interface{}) string {
	strSlice := []string{}
	v := reflect.ValueOf(ivalue)
	for _, key := range v.MapKeys() {
		strSlice = append(strSlice, fmt.Sprintf("%s:%s", key, v.MapIndex(key)))
	}
	return quotedOption(name, strings.Join(strSlice, ","))
}

func optionFromSlice(name string, ivalue interface{}) string {
	strSlice := []string{}
	v := reflect.ValueOf(ivalue)
	for i := 0; i < v.Len(); i++ {
		strSlice = append(strSlice, fmt.Sprint(v.Index(i)))
	}

	if name == "preview" {
		return quotedOption(name, strings.Join(strSlice, " "))
	} else {
		return quotedOption(name, strings.Join(strSlice, ","))
	}
}

func optionFromBool(name string, ivalue interface{}) string {
	if ivalue == false {
		return fmt.Sprintf("--no-%s", name)
	}
	return fmt.Sprintf("--%s", name)
}

func optionFromNumber(name string, ivalue interface{}) string {
	return fmt.Sprintf("--%s=%v", name, ivalue)
}

func optionFromString(name string, ivalue interface{}) string {
	return quotedOption(name, ivalue)
}

func quotedOption(key string, ivalue interface{}) string {
	if strings.Contains(ivalue.(string), " ") {
		return fmt.Sprintf("--%s=%q", key, ivalue)
	} else {
		return fmt.Sprintf("--%s=%s", key, ivalue)
	}
}
