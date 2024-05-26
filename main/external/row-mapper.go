package external

import "database/sql"

type RowsMapper interface {
	MapRows(row *sql.Rows) (interface{}, error)
}
