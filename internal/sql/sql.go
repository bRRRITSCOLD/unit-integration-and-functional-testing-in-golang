package sequel

import (
	"errors"
	"fmt"
	"unit-integration-and-functional-testing-in-golang/internal/clients"

	_ "github.com/go-sql-driver/mysql"
)

const (
	GetUserQuery = "SELECT id, email FROM users WHERE id=%d;"
)

var (
	dbClient *clients.SQLClientInterface
)

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error

	dbClient, err = clients.SQLClientInterface.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "test_user", "123abc", "127.0.0.1:3306", "test"))
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
