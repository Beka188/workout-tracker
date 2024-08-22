package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, pool *pgxpool.Pool, email, username, password string) (int, error) {
	var userID int

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	query := `
        INSERT INTO users (email, username, password)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	err = pool.QueryRow(ctx, query, email, username, passHash).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return userID, nil
}

func UpdateUser(ctx context.Context, pool *pgxpool.Pool, userID int, password string) error {
	_, err := GetUser(ctx, pool, userID)
	if err != nil {
		return err
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `
        UPDATE users
        SET password = $1
        WHERE id = $2
    `

	_, err = pool.Exec(ctx, query, passHash, userID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// func DeleteUser(ctx context.Context, pool *pgxpool.Pool, userID int) error {

// }

func GetUser(ctx context.Context, pool *pgxpool.Pool, userID int) (User, error) {
	// Define the SQL query to get the user
	query := `
        SELECT email, username
        FROM users
        WHERE id = $1
    `
	var user User
	err := pool.QueryRow(ctx, query, userID).Scan(&user.Email, &user.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return User{}, fmt.Errorf("user with id %d not found", userID)
		}
		fmt.Println(err)
		return User{}, err
	}

	return user, nil
}

func All(ctx context.Context, pool *pgxpool.Pool) ([]User, error) {
	var users []User

	// SQL query to select all users
	query := `
        SELECT id, username, password
        FROM users
    `

	// Execute the query
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into User structs
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
