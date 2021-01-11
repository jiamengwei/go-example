package main

import "net/http"

func users() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request.Header.Set("content-type", "application/json; charset=utf-8")
		writer.Write([]byte("Jack"))
		writer.WriteHeader(200)
	})
}

func main() {
	//启动一个文件服务器
	http.Handle("/", http.FileServer(http.Dir("D:/")))
	http.Handle("/users", users())
	http.ListenAndServe(":9999", nil)
}
