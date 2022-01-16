package main

import (
	"fmt"
	"go_lect/internal/db"
	"go_lect/internal/db/repository"
	"go_lect/internal/services"
	"log"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userService := services.NewUserService(conn)

	user := repository.User{
		Name:         "test user",
		Email:        "test.user@gmail.com",
		PasswordHash: "somepasshash",
	}
	address := repository.UserAddress{
		Country: "Ukraine",
		City:    "Kharkiv",
		Address: "Test street",
	}
	err = userService.CreateUser(user, address)
	fmt.Println(err)

}
