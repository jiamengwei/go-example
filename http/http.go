package main

import "net/http"

func main() {
	//启动一个文件服务器
	http.Handle("/", http.FileServer(http.Dir("D:/")))

	http.ListenAndServe(":9999", nil)
}
