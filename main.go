package main

import (
	"github.com/ZNemuZ/outly-back/controller"
	"github.com/ZNemuZ/outly-back/db"
	"github.com/ZNemuZ/outly-back/repository"
	"github.com/ZNemuZ/outly-back/router"
	"github.com/ZNemuZ/outly-back/usecase"
	"github.com/ZNemuZ/outly-back/validator"
)

func main() {
	db := db.NewDb()
	userValidator := validator.NewUserValidator()
	postValidator := validator.NewPostValidator()
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepository, postValidator)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	e := router.NewRouter(userController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
