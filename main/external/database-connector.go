package external

import (
	"database/sql"
	"log"
)

type DatabaseConnector struct {
	ConnectionString string
	Database         *sql.DB
}

func NewDatabaseConnector() *DatabaseConnector {
	return &DatabaseConnector{ConnectionString: "postgres://postgres:postgres@localhost/poketracker?sslmode=disable"}
}

func (i *DatabaseConnector) Query(query string, mapper RowsMapper, args ...any) (interface{}, error) {
	err := i.connect()
	if err != nil {
		return nil, err
	}

	log.Printf("Executing query: %v\n", query)
	log.Printf("Args: %v\n", args)
	rows, err := i.Database.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)
	defer func(i *DatabaseConnector) {
		err = i.close()
	}(i)

	mapRows, err := mapper.MapRows(rows)
	return mapRows, err
}

func (i *DatabaseConnector) Execute(query string, args ...any) (int, error) {
	err := i.connect()
	if err != nil {
		return 0, err
	}

	log.Printf("Executing query: %v\n", query)
	res, err := i.Database.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	count, err := i.getUpdatedRowCount(res)
	if err != nil {
		return 0, err
	}

	err = i.close()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (i *DatabaseConnector) connect() error {
	db, err := sql.Open("postgres", i.ConnectionString)
	i.Database = db
	return err
}

func (i *DatabaseConnector) close() error {
	return i.Database.Close()
}

func (i *DatabaseConnector) getUpdatedRowCount(result sql.Result) (int, error) {
	affected, err := result.RowsAffected()
	log.Printf("Rows affected: %v", affected)
	return int(affected), err
}
