package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var logDir = "E:\\apprun\\littlec-log\\smartbuild\\"
var outDir = "E:\\apprun\\base64\\"

func main() {

	//file, _ := os.Open("E:\\git-space\\go-example\\io\\log\\test.txt")
	//reader := bufio.NewReader(file)
	//for{
	//	line, prefix, err := reader.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(prefix)
	//	fmt.Println(string(line))
	//}

	dir, err := ioutil.ReadDir(logDir)
	if err != nil {
		log.Fatal("文件夹读取失败：", err)
	}

	//readLog("default.log")

	c := make(chan bool)
	for _, v := range dir {
		if !strings.HasPrefix(v.Name(), "default") {
			continue
		}
		fmt.Println("开始解析文件：", v.Name())
		go chanWrap(v.Name(), c)
		<-c
	}
}

func chanWrap(filename string, c chan bool) {
	readLog(filename)
	c <- true
}

func readLog(filename string) {
	fmt.Println(filename)
	file, err := os.Open(logDir + filename)
	if err != nil {
		log.Println("文件读取失败", err)
	}

	reader := bufio.NewReader(file)
	name, phone := "", ""
	for {
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		text := string(line)
		if !prefix && !strings.Contains(text, "<200 OK,GETFaceUserApiResponse(bean=GETFaceUserApiResponse.BeanDTO(errmsg=OK,") {
			continue
		}

		if prefix && strings.Contains(text, "<200 OK,GETFaceUserApiResponse(bean=GETFaceUserApiResponse.BeanDTO(errmsg=OK,") {
			//fmt.Println(text)
			phone = strings.Split(strings.Split(text, "mobile=")[1], ",")[0]
			name = strings.Split(strings.Split(text, "userName=")[1], ",")[0]
			base64Img := strings.Split(strings.Split(text, "facePic=")[1], ")])),")[0]
			//fmt.Println(strings.Split(text, "facePic=")[1])
			decodeString, _ := base64.StdEncoding.DecodeString(strings.Replace(base64Img, " ", "\n", -1))
			ioutil.WriteFile(outDir+"\\image\\"+name+".jpg", decodeString, 0666)
			err := ioutil.WriteFile(outDir+name+"_"+phone+".txt", []byte(base64Img), 0777)
			if err != nil {
				log.Println("写入文件失败", err)
			}
			continue
		}

		openFile, err := os.OpenFile(outDir+name+"_"+phone+".txt", os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			log.Println("追加文件失败", err)
		}
		openFile.WriteString(text)
		openFile.Close()
	}
}
