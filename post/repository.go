package post

import (
	"database/sql"
	"go.example/blog/db"
	"log"
	"time"
)

func QueryAll(qTitle string) []post {
	var rows *sql.Rows
	var err error
	if qTitle != "" {
		rows, err = db.Conn().Query("select * from post where title like ? order by create_time desc", "%"+qTitle+"%")
	} else {
		rows, err = db.Conn().Query("select * from post order by create_time desc")
	}
	defer rows.Close()
	if err != nil {
		log.Fatal("blog query failed", err)
	}
	var posts []post

	for rows.Next() {
		var id, categoryId int
		var title, content string
		var createTime, updateTime time.Time
		err := rows.Scan(&id, &title, &content, &createTime, &updateTime, &categoryId)
		if err != nil {
			log.Fatal("blog rows scan failed", err)
		}
		posts = append(posts, New(id, title, content, createTime, updateTime, categoryId))
	}

	return posts
}

func QueryByCategory(qCategoryName string) []post {
	rows, err := db.Conn().Query("select post.* from post inner join category on post.category_id = category.id where category.name = ?", qCategoryName)
	if err != nil {
		log.Println("post.QueryByCategory 查询失败", err)
		return nil
	}
	var posts []post

	for rows.Next() {
		var id, category_id int
		var title, content string
		var createTime, updateTime time.Time
		rows.Scan(&id, &title, &content, &createTime, &updateTime, &category_id)
		posts = append(posts, New(id, title, content, createTime, updateTime, category_id))
	}

	return posts
}

func QueryById(qId int) *post {
	rows, err := db.Conn().Query("select * from post where id = ?", qId)
	if err != nil {
		log.Println(err)
		return nil
	}
	existing := rows.Next()
	if !existing {
		return nil
	}
	var id, categoryId int
	var title, content string
	var createTime, updateTime time.Time
	err = rows.Scan(&id, &categoryId, &title, &content, &createTime, &updateTime)
	if err != nil {
		log.Println(err)
		return nil
	}
	p := New(id, title, content, createTime, updateTime, categoryId)
	return &p
}

func QueryByTitle(qTitle string) *post {
	rows, err := db.Conn().Query("select * from post where title = ?", qTitle)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	var id, categoryId int
	var title, content string
	var createTime, updateTime time.Time
	existing := rows.Next()
	if !existing {
		return nil
	}

	err = rows.Scan(&id, &title, &content, &createTime, &updateTime, &categoryId)
	if err != nil {
		log.Println(err)
		return nil
	}

	p := New(id, title, content, createTime, updateTime, categoryId)
	return &p
}

func UpdatePost(newPost post) int64 {
	prepare, err := db.Conn().Prepare("update post set title = ?, content = ? where id = ?")
	defer prepare.Close()
	if err != nil {
		log.Println(err)
		return 0
	}
	exec, err := prepare.Exec(newPost.Title, newPost.Content, newPost.Id)
	if err != nil {
		log.Println(err)
		return 0
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return affected
}

func SavePost(newPost post) (int64, error) {
	prepare, err := db.Conn().Prepare("insert into post (title, content,category_id) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	exec, err := prepare.Exec(newPost.Title, newPost.Content, newPost.CategoryId)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return exec.RowsAffected()
}

func DeleteById(id int) int64 {
	exec, err := db.Conn().Exec("delete from post where id = ?", id)
	if err != nil {
		log.Println(err)
		return 0
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	return affected
}
