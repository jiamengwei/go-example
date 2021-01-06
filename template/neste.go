package main

import (
	"fmt"
	"html/template"
	"os"
)

func nested() {
	s1, _ := template.ParseFiles("template/header.html", "template/content.html", "template/footer.html")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", "content")
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
