package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func readDir(dirname string) {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		fmt.Printf("fileName:%s, modTime:%s, size:%d, isDis:%t, mode:%s \n", file.Name(), file.ModTime(), file.Size(), file.IsDir(), file.Mode())
	}
}

func readFile(filename string) (content string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if utf8.Valid(file) {
		return string(file)
	}
	reader := transform.NewReader(bytes.NewReader(file), simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return string(all)
}

func tempFile(dir, pattern string) (filename string) {
	file, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	return file.Name()
}

func writerFile(filename string, content []byte) {
	err := ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	readDir("D://")
	content := readFile("D://getDevicesSuccess.json")
	fmt.Println(content)
	filename := tempFile("D://tmp", "*.temp.file")
	os.RemoveAll(filename)
	writerFile("D://tmp/hello2.txt", []byte("hello12"))
}
