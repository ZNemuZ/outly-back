package router

import (
	"os"

	"github.com/ZNemuZ/outly-back/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	p := e.Group("/posts")
	//postsに関連するエンドポイントに向かう際に必ず実行される
	p.Use(echojwt.WithConfig(echojwt.Config{ //認証ミドルウェアに設定を渡す
		SigningKey:  []byte(os.Getenv("SECRET")), //jwtの署名を秘密鍵で確認
		TokenLookup: "cookie:token",              //リクエスト内のtokenの位置の指定
	}))
	p.GET("", pc.GetAllPosts)
	p.GET(("/:postId"), pc.GetPostById)
	p.POST("", pc.CreatePost)
	p.DELETE("/:postId", pc.DeletePost)
	return e
}
