package router

import (
	"net/http"
	"os"

	"github.com/ZNemuZ/outly-back/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", os.Getenv("FE_URL")}, //許可するオリジンの設定
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken}, //許可するヘッダー
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true, //認証情報を含めるリクエストを許可する
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"), //cookieの送信される範囲を指定
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //postMan用
		// CookieMaxAge: 60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
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
