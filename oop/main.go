package main

import (
	"example.com/hello/oop/employee"
	"example.com/hello/oop/post"
)

func main() {
	e := employee.New("jia", 23)
	e.Info()

	p := post.NewPost("My Blog", post.NewArticle("My First Blog", "Haha hello"))
	p.Info()
}
