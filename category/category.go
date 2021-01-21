package category

import "time"

type category struct {
	id          int
	Name        string `json:"name"`
	Description string `json:"description" binding:"required"`
	CreateTime  time.Time
}

func New(id int, name string, description string, createTime time.Time) *category {
	return &category{
		id:          id,
		Name:        name,
		Description: description,
		CreateTime:  createTime,
	}
}
