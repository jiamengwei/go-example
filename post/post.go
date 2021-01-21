package post

import "time"

type post struct {
	Id         int
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	CategoryId int
}

func New(id int, title, content string, createTime, updateTime time.Time, categoryId int) post {
	return post{
		Id:         id,
		Title:      title,
		Content:    content,
		CreateTime: createTime,
		UpdateTime: updateTime,
		CategoryId: categoryId,
	}
}
