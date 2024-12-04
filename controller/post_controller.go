package controller

import (
	"net/http"
	"strconv"

	"github.com/ZNemuZ/outly-back/model"
	"github.com/ZNemuZ/outly-back/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAllPosts(c echo.Context) error
	GetPostById(c echo.Context) error
	CreatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}
type postController struct {
	pu usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	//トークンをjwtToken型として取得
	user := c.Get("user").(*jwt.Token)
	//ペイロード部分をMapClaimsと取得
	claims := user.Claims.(jwt.MapClaims)
	//ペイロードからuserIdを取得
	userId := claims["user_id"]
	postsRes, err := pc.pu.GetAllPosts(uint(userId.(float64))) //※jwtのペイロードは数値がfloatで扱われるためfloat64で型アサーションをしている
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postsRes)
}

func (pc *postController) GetPostById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	//URLのpostIdに対応する値を取得
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)
	postRes, err := pc.pu.GetPostById(uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	post.UserId = uint(userId.(float64))
	postRes, err := pc.pu.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) DeletePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	err := pc.pu.DeletePost(uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
