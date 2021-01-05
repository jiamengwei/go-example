package main

import (
	"fmt"
	"html/template"
	"os"
)

func nested() {
	s1, _ := template.ParseFiles("template/header.html", "template/content.html", "template/footer.html")
	//s1.ExecuteTemplate(os.Stdout, "header", "header")
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", map[string]string{"header": "h", "content": "cc", "footer": "fff"})
	fmt.Println()
	//s1.ExecuteTemplate(os.Stdout, "footer", "footer")
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
