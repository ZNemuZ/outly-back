package main

import (
	"github.com/ZNemuZ/outly-back/controller"
	"github.com/ZNemuZ/outly-back/db"
	"github.com/ZNemuZ/outly-back/repository"
	"github.com/ZNemuZ/outly-back/router"
	"github.com/ZNemuZ/outly-back/usecase"
)

func main() {
	db := db.NewDb()
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	e := router.NewRouter(userController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
