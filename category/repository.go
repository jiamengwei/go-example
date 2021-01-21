package category

import (
	"errors"
	"go.example/blog/db"
	"log"
	"time"
)

func saveCategory(category category) (int64, error) {
	prepare, err := db.Conn().Prepare("insert into category (name, description) values (?,?)")
	if err != nil {
		log.Fatal("分类保存失败：", err)
		return 0, errors.New("分类保存失败")
	}
	exec, err := prepare.Exec(category.Name, category.Description)
	return exec.LastInsertId()
}

func QueryById(qId int) *category {
	row := db.Conn().QueryRow("select * from category where id = ?", qId)
	var id int
	var categoryName, description string
	var createTime time.Time
	err := row.Scan(&id, &categoryName, &description, &createTime)
	if err != nil {
		log.Println("分类查询失败 ", err)
		return nil
	}
	return New(id, categoryName, description, createTime)
}

func queryByName(qName string) *category {
	row := db.Conn().QueryRow("select * from category where name = ?", qName)
	var id int
	var name, description string
	var createTime time.Time
	err := row.Scan(&id, &name, &description, &createTime)
	if err != nil {
		log.Println("分类查询失败 ", err)
		return nil
	}
	return New(id, name, description, createTime)
}

func deleteById(id string) (int64, error) {
	result, err := db.Conn().Exec("delete from category where id = ?", id)
	if err != nil {
		log.Println("分类删除失败", err)
		return 0, errors.New("分类删除失败")
	}
	return result.RowsAffected()
}
