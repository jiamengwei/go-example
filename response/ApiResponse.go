package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiResponse struct {
	Msg  string
	Code int
	Data interface{}
}

func (res apiResponse) Write(writer http.ResponseWriter) {
	jsonResult, err := json.Marshal(res)
	if err != nil {
		log.Panic("json解析失败：{}", err)
	}
	writer.Header().Set("content-type", "text/json")
	writer.Write(jsonResult)
	writer.WriteHeader(200)
}

func Fail(msg string) apiResponse {
	return apiResponse{
		Code: 1,
		Msg:  msg,
		Data: nil,
	}
}

func Success(msg string, data interface{}) apiResponse {
	return apiResponse{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}

func EmptySlice() []interface{} {
	return make([]interface{}, 0)
}
