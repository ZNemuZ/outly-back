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
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
