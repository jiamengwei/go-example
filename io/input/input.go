package main

import (
	"bufio"
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type msg struct {
	Content string `json:"content"`
}

func (m *msg) toJson() string {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		log.Panic(err)
	}
	return string(jsonBytes)
}
func push(body string) {

	msg := &msg{
		Content: body,
	}
	content := msg.toJson()
	res, err := http.Post("", "application/json", strings.NewReader(content))
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
}

func editMode() {
	scanner := bufio.NewScanner(os.Stdin)
	contentBuffer := make([]string, 0, 100)
	needPush := false
	for scanner.Scan() {
		s := scanner.Text()
		if s == "push" {
			needPush = true
			break
		}

		if s == "exit" {
			break
		}
		contentBuffer = append(contentBuffer, s, "\n")
	}
	if needPush {
		content := strings.Join(contentBuffer, "")
		push(content)
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		push(strings.Join(args[1:], " "))
	} else {
		editMode()
	}

}
