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
		result = append(result, *row)
	}
	return result, nil
}

func (i *UserMapper) mapRow(row *sql.Rows) (*domain.User, error) {
	s := new(domain.User)
	err := row.Scan(&s.UserId, &s.AvatarUrl, &s.BulkMode)
	if err != nil {
		return s, err
	}
	return s, nil
}
