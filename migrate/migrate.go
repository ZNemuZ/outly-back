package main

import (
	"fmt"

	"github.com/ZNemuZ/outly-back/db"
	"github.com/ZNemuZ/outly-back/model"
)

func main() {
	dbConn := db.NewDb()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Post{})
}
