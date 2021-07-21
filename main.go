package main

//利用するパッケージの宣言
import (
	"net/http"
	"strconv"
	"time"

	//外部パッケージ
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"

//グローバル変数eにcreateMux()の関数の戻り値を格納
var e = createMux()

func main() {
	//`/`と言うパス(URL)と``articleindex`を結びつける
	e.GET("/", articleindex)
	//ルーディング追加
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow) //idはパスパラメーター
	e.GET("/:id/edit", articleEdit)

	//wevサーバーをローカルホストで起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	e.Static("js", "src/js")

	return e
}

//データ渡し
func articleindex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article index",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}
	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	//パスパラメーターを抽出
	id, _ := strconv.Atoi(c.Param("id")) //c.Param()文字列型からAtoi()でキャスト変換

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

//生成したhtmlデータをバイトデータとして返す
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	//htmlblob()からhtmlをバイトデータとして受け取る
	b, err := htmlBlob(file, data)
	//エラーチェック
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	//ステータスコード200でhtmlデータをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
