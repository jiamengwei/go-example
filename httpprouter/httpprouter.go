package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type user struct {
	Name string
}

type response struct {
	Code bool
	Msg  string
	Data user
}

func Json(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res response
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		fmt.Println("decode fail")
	}
	w.Header().Set("content-type", "text/json")
	fmt.Printf("%+v", res)
	marshal, err := json.Marshal(res)
	w.Write(marshal)
}

func main() {
	//res := response{
	//	Code: false,
	//	Msg:  "hello",
	//	Data: user{
	//		Name: "jack",
	//	},
	//}
	//jsonStr, err := json.Marshal(res)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(string(jsonStr))
	//}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/json", Json)
	log.Fatal(http.ListenAndServe(":9090", router))
}
