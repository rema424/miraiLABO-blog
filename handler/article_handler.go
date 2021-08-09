package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"miraiLABO-blog/model"
	"miraiLABO-blog/repository"

	"github.com/labstack/echo/v4"
)

func Articleindex(c echo.Context) error {
	if c.Request().URL.Path == "/articles" {
		c.Redirect(http.StatusPermanentRedirect, "/")
	}

	articles, err := repository.ArticleListByCursor(0)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	var cursor int
	if len(articles) != 0 {
		cursor = articles[len(articles)-1].ID
	}

	data := map[string]interface{}{
		"Articles": articles,
		"Cursor":   cursor,
	}
	return render(c, "article/index.html", data)
}

func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := repository.ArticleGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Article": article,
	}
	return render(c, "article/show.html", data)
}

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}
	return render(c, "article/edit.html", data)
}

// ArticleCreateOutput ...
type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleCreate ...
func ArticleCreate(c echo.Context) error {
	// 送信されてくるフォームの内容を格納する構造体を宣言します。
	var article model.Article

	// レスポンスとして返却する構造体を宣言します。
	var out ArticleCreateOutput

	// フォームの内容を構造体に埋め込みます。
	if err := c.Bind(&article); err != nil {
		// エラーの内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// リクエストの解釈に失敗した場合は 400 エラーを返却します。
		return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェックを実行します。
	if err := c.Validate(&article); err != nil {
		// エラーの内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// エラーの内容を検査してカスタムエラーメッセージを取得します。
		out.ValidationErrors = article.ValidationErrors(err)

		// 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
		return c.JSON(http.StatusUnprocessableEntity, out)

	}

	// repository を呼び出して保存処理を実行します。
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		// エラーの内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, out)
	}

	// SQL 実行結果から作成されたレコードの ID を取得します。
	id, _ := res.LastInsertId()

	// 構造体に ID をセットします。
	article.ID = int(id)

	// レスポンスの構造体に保存した記事のデータを格納します。
	out.Article = &article

	// 処理成功時はステータスコード 200 でレスポンスを返却します。
	return c.JSON(http.StatusOK, out)
}

// ArticleDelete...
func ArticleDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))
	if err := repository.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}

// ArticleList...
func ArticleList(c echo.Context) error {
	cursor, _ := strconv.Atoi(c.QueryParam("cursor"))
	articles, err := repository.ArticleListByCursor(cursor)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, articles)
}
