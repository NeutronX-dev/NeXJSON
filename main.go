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

func Stringify(JSON map[string]interface{}) (string, error) {
	ret, err := json.Marshal(JSON)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func Access(JSON interface{}, props string /* Value.MyOtherValue */) interface{} {
	var properties []string = strings.Split(props, ".")

	var Last interface{} = JSON
	for _, property := range properties {
		Last = (Last.(map[string]interface{}))[property]
	}
	return Last
}

func AllKeys(JSON interface{}) []string {
	mapped := JSON.(map[string]interface{})
	res := make([]string, 0)
	for key := range mapped {
		res = append(res, key)
	}
	return res
}
