package external

import (
	"database/sql"
	"errors"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (i *UserMapper) MapRows(rows *sql.Rows) (interface{}, error) {
	if rows.Next() {
		return i.mapRow(rows)
	} else {
		return nil, errors.New("user not found")
	}
}

func (i *UserMapper) mapRow(row *sql.Rows) (int, error) {
	userId := 0
	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
