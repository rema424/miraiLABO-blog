package main

//利用するパッケージの宣言
import (
	"net/http" //標準パッケージ

	//外部パッケージ
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//グローバル変数eにcreateMux()の関数の戻り値を格納
var e = createMux()

func main() {
	//`/`と言うパス(URL)と``articleindex`を結びつける
	e.GET("/", articleindex)

	//wevサーバーをローカルホストで起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e //アプリケーションインスタンスを返却
}

func articleindex(c echo.Context) error {
	//ステータスコード200で、"Hello, World!"と言う文字列をレスポンス
	return c.String(http.StatusOK, "Hello, World!")
}
