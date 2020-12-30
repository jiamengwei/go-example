package post

import "fmt"

type post struct {
	name string
	article
}

func NewPost(name string, a article) post {
	return post{
		name:    name,
		article: a,
	}
}

func (p post) Info() {
	fmt.Printf("post name: %s, %s\n", p.name, p.article.Info())
}

type article struct {
	title   string
	content string
}

func (a article) Info() string {
	return fmt.Sprintf("title:%s, content:%s \n", a.content, a.title)
}

func NewArticle(title, content string) article {
	return article{
		title:   title,
		content: content,
	}
}
