package external

import (
	"database/sql"
)

type DatabaseConnector struct {
	ConnectionString string
	Database         *sql.DB
}

func NewDatabaseConnector() *DatabaseConnector {
	return &DatabaseConnector{ConnectionString: "postgres://postgres:postgres@localhost/poketracker?sslmode=disable"}
}

func (i *DatabaseConnector) Query(query string, mapper RowsMapper, args ...any) interface{} {
	i.connect()

	rows, err := i.Database.Query(query, args...)
	if err != nil {
		panic(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	defer i.close()

	return mapper.MapRows(rows)
}

func (i *DatabaseConnector) connect() {
	db, err := sql.Open("postgres", i.ConnectionString)
	if err != nil {
		panic(err)
	}
	i.Database = db
}

func (i *DatabaseConnector) close() {
	err := i.Database.Close()
	if err != nil {
		panic(err)
	}
}
