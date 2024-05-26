package external

import (
	"database/sql"
)

type EditionMapper struct {
}

func NewEditionMapper() *EditionMapper {
	return &EditionMapper{}
}

func (i *EditionMapper) MapRows(rows *sql.Rows) (interface{}, error) {
	result := make([]string, 0)
	for rows.Next() {
		row, err := i.mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (i *EditionMapper) mapRow(row *sql.Rows) (string, error) {
	edition := ""
	err := row.Scan(&edition)
	if err != nil {
		return "", err
	}
	return edition, nil
}
