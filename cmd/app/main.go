package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"workout-tracker/internal/db"
	"workout-tracker/internal/services"
	"workout-tracker/internal/user"
)

func main() {
	token, err := services.CreateToken("email12", 12)
	fmt.Println(token)
	ok, err := services.VerifyToken(token, "email13", 12)
	fmt.Println(ok)
	return
	pool, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.CloseDB()

	userID, err := user.Register(context.Background(), pool, "email", "newuser", "hashed_password123")
	if err != nil {
		log.Fatal("Failed to create user:", err)
	}

	// users, err := user.All(context.Background(), pool)
	// if err != nil {
	// 	log.Fatal("Failed to get all users:", err)
	// }
	// fmt.Println(users)
	log.Printf("User created successfully with ID %d", userID)
}
