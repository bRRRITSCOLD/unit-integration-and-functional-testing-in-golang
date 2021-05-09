package clients

import (
	"database/sql"
	"errors"
)

type row struct{}

type SQLClient struct {
	db *sql.DB
}

type SQLClientInterface interface {
	Query(query string, args ...interface{}) (*row, error)
}

func Open(driverName, dataSourceName string) (SQLClientInterface, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	// db, err := SQLClientInterface.Open(driverName, dataSourceName)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	client := SQLClient{
		db: db,
	}
	return client, nil
}

func (sC SQLClient) Query(query string, args ...interface{}) (*row, error) {
	return nil, nil
}
