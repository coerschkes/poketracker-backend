package external

import (
	"database/sql"
	"poketracker-backend/main/domain"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (i *UserMapper) MapRows(rows *sql.Rows) (interface{}, error) {
	result := make([]domain.User, 0)
	for rows.Next() {
		row, err := i.mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.User{UserId: row[0], AvatarUrl: row[1]})
	}
	return result, nil
}

func (i *UserMapper) mapRow(row *sql.Rows) ([]string, error) {
	s := make([]string, 2)
	err := row.Scan(&s[0], &s[1])
	if err != nil {
		return s, err
	}
	return s, nil
}
