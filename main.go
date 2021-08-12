package main

//利用するパッケージの宣言
import (
	"log"
	"os"

	"miraiLABO-blog/handler"
	"miraiLABO-blog/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

//グローバル変数eにcreateMux()の関数の戻り値を格納
var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	//ルーディングのグループ
	auth := e.Group("")

	auth.Use(basicAuth())

	// TOPページに一覧を表示
	e.GET("/", handler.Articleindex)

	e.GET("/articles", handler.Articleindex) //一覧画面
	//e.GET("/articles/new", handler.ArticleNew)              // 新規作成画面
	auth.GET("/articles/new", handler.ArticleNew)      // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow) // 詳細画面
	//e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面
	auth.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	e.GET("/api/articles", handler.ArticleList) // 一覧
	//e.POST("/api/articles", handler.ArticleCreate)              // 作成
	auth.POST("/api/articles", handler.ArticleCreate) // 作成
	//e.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 消去
	auth.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 消去
	//e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)  // 更新
	auth.PATCH("/api/articles/:articleID", handler.ArticleUpdate) // 更新

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New() //アプリケーションインスタンスを生成

	//アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())
	//e.Use(basicAuth())

	e.Static("/css", "src/css")
	e.Static("js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

// CustomValidator...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func basicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "SteinsGate" && password == "ElPsyCongroo" {
			return true, nil
		}
		return false, nil
	})
}
