package repository

import (
	"miraiLABO-blog/model"
)

// ArticleList...
func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	var articles []*model.Article
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	return articles, nil
}
