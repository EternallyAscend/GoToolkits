package JSON

import (
	"encoding/json"
	"errors"
	"fmt"
)

func TransferStringToJson(stringData string) (interface{}, error) {
	var interfaces interface{}
	var bytes []byte
	bytes = []byte(stringData)
	err := json.Unmarshal(bytes, &interfaces)
	if nil != err {
		return nil, err
	}
	return interfaces, nil
}

func TransferJsonToString() {

}

func ListItemInJson(jsonString string) (string, error) {
	var interfaces interface{}
	interfaces, err := TransferStringToJson(jsonString)
	if nil != err {
		return "", err
	}
	content, ok := interfaces.(map[string]interface{})
	if ok {
		for k, v := range content {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "Unknown Type.")
			}
		}
	}
	return "", nil
}

func SearchItemInJson(jsonString string, keyword string) (string, error) {
	var interfaces interface{}
	interfaces, err := TransferStringToJson(jsonString)
	if nil != err {
		return "", err
	}
	content, ok := interfaces.(map[string]interface{})
	if ok {
		for k, v := range content {
			if k != keyword {
				continue
			}
			switch v2 := v.(type) {
			case string:
				return v2, nil
			case int:
				return string(v2), nil
			case bool:
				if v2 {
					return "true", nil
				} else {
					return "false", nil
				}
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
				return "", errors.New("Result is an array. ")
			default:
				return "", errors.New("Unknown Type. ")
			}
		}
	}
	return "", errors.New("Cannot found value. ")
}

func GetAllItemsInJson(jsonString string) ([]string, []string, error) {
	var key []string
	var value []string
	key = []string{}
	value = []string{}
	var interfaces interface{}
	interfaces, err := TransferStringToJson(jsonString)
	if nil != err {
		return nil, nil, err
	}
	content, ok := interfaces.(map[string]interface{})
	if ok {
		for k, v := range content {
			key = append(key, k)
			switch v2 := v.(type) {
			case string:
				value = append(value, v2)
			case int:
				value = append(value, string(v2))
			case bool:
				if v2 {
					value = append(value, "true")
				} else {
					value = append(value, "false")
				}
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
				return nil, nil, errors.New("Result is an array. ")
			default:
				return nil, nil, errors.New("Unknown Type. ")
			}
		}
	}
	return key, value, errors.New("Cannot found value. ")
}

// Another solution with unique structure.

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func TestJson() {
	var s ServerSlice

	str := `{
                "servers": [
                    {
                        "serverName": "Shanghai",
                        "serverIP": "127.0.0.1"
                    }, {
                        "serverName": "Beijing",
                        "serverIP": "127.0.0.2"
                    }
                ]
            }`

	json.Unmarshal([]byte(str), &s)

	for key, val := range s.Servers {

		print(`Key：`, key, "\t")
		print(`Name：`, val.ServerName, "\t")
		println(`IP：`, val.ServerIP)
	}
}
