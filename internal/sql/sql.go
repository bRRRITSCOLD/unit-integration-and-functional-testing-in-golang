package sql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	GetUserQuery = "SELECT id, email FROM users WHERE id=%d;"
)

var (
	dbClient *sql.DB
)

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	dbClient, err = sql.Open("mysql", "this is the connection string")
	if err != nil {
		panic(err)
	}
}

func GetUser(userId int64) (*User, error) {
	rows, err := dbClient.Query(fmt.Sprintf(GetUserQuery, userId))
	if err != nil {
		panic(err)
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
