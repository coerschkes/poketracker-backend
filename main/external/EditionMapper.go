package external

import (
	"database/sql"
)

type EditionMapper struct {
}

func NewEditionMapper() *EditionMapper {
	return &EditionMapper{}
}

func (i *EditionMapper) MapRows(rows *sql.Rows) interface{} {
	result := make([]string, 0)
	for rows.Next() {
		result = append(result, i.mapRow(rows))
	}
	return result
}

func (i *EditionMapper) mapRow(row *sql.Rows) string {
	edition := ""
	err := row.Scan(&edition)
	if err != nil {
		panic(err)
	}
	return edition
}
