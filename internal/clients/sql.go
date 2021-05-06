package clients

import (
	"database/sql"
	"errors"
)

type row struct{}

type sqlClient struct {
	db *sql.DB
}

type sqlClientInterface interface {
	Query(query string, args ...interface{}) (*row, error)
}

func Open(driverName, dataSourceName string) (sqlClientInterface, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	client := sqlClient{
		db: db,
	}
	return client, nil
}

func (sC sqlClient) Query(query string, args ...interface{}) (*row, error) {

}
