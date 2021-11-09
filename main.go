package NeXJSON

import (
	"encoding/json"
	"strings"
)

func Parse(str string) (interface{}, error) {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(str), &payload)
	if err != nil {
		return (nil), err
	}
	return payload, nil
}

func Stringify(data_ interface{}) (string, error) {
	ret, err := json.Marshal(data_)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func Access(data_ interface{}, props string /* Value.MyOtherValue */) interface{} {
	var properties []string = strings.Split(props, ".")

	var Last interface{} = data_
	for _, property := range properties {
		Last = (Last.(map[string]interface{}))[property]
	}
	return Last
}

func Replace(data_ interface{}, props string, NewValue interface{}) {
	var properties []string = strings.Split(props, ".")
	var Last interface{} = data_
	for i, property := range properties {
		if i == len(properties)-1 {
			(Last.(map[string]interface{}))[property] = NewValue
			break
		}
		Last = (Last.(map[string]interface{}))[property]
	}
}

func AllKeys(data_ interface{}) []string {
	mapped := data_.(map[string]interface{})
	res := make([]string, 0)
	for key := range mapped {
		res = append(res, key)
	}
	return res
}
