package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Article struct {
	ID      int       `db:"id" form:"id" json:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50" json:"title"`
	Body    string    `db:"body" form:"body" validate:"required" json:"body"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

func (a *Article) ValidationErrors(err error) []string {
	var errMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "Title":
			//エラーになったバリデーションルールを特定します。
			switch err.Tag() {
			case "required":
				message = "タイトルは必須です。"
			case "max":
				message = "タイトルは最大50文字です。"
			}
		case "Body":
			message = "本文は必須です。"
		}

		// メッセージをスライスに追加する。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}
