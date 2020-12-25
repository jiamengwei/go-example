package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ApiResponse struct {
	Code int
	Msg  string
	Data interface{}
}

func parseJSON(jsonContent string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonContent), obj)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseJSONToMap(jsonContent string) (jsonMap map[string]interface{}) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonContent), &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	m := jsonData.(map[string]interface{})
	return m
}

func convertToJSON(obj interface{}) (jsonContent string) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

func main() {

	res := ApiResponse{
		Code: 200,
		Msg:  "OK",
		Data: "你好 ！",
	}

	jsonContent := convertToJSON(res)
	fmt.Println("convert res to json :", jsonContent)

	var newRes ApiResponse
	parseJSON(jsonContent, &newRes)
	fmt.Println(newRes)

	jsonMap := parseJSONToMap(jsonContent)
	for k, v := range jsonMap {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string ", vv)
		case int:
			fmt.Println(k, "is int ", vv)
		default:
			fmt.Printf("%s is %f, type is %T \n", k, vv, vv)
		}
	}

}
