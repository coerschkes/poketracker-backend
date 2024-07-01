package external

import (
	"errors"
	"fmt"
	"log"
	"poketracker-backend/main/domain"
)

const (
	selectUserQuery = "SELECT avatarUrl, bulkmode FROM userinfo WHERE userinfo.userId = $1"
	createUserQuery = "INSERT INTO userinfo (userId, avatarurl, bulkmode) VALUES ($1, $2, $3)"
	updateUserQuery = "UPDATE userinfo SET avatarurl = $2, bulkmode = $3 WHERE userinfo.userId = $1"
	deleteUserQuery = "DELETE FROM userinfo WHERE userinfo.userId = $1"
)

type UserRepository interface {
	Find(userId string) (interface{}, error)
	Create(userId string, user *domain.User) error
	Update(userId string, user *domain.User) error
	Delete(userId string) error
}

type UserRepositoryImpl struct {
	connector *DatabaseConnector
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{connector: NewDatabaseConnector()}
}

func (p UserRepositoryImpl) Find(userId string) (interface{}, error) {
	query, err := p.connector.Query(selectUserQuery, NewUserMapper(), userId)
	if err != nil {
		log.Printf("user-repository.Find(): error while fetching user: %v\n", err)
		return nil, errors.New("error while fetching user")
	}

	users := query.([]domain.User)
	if len(users) > 1 {
		log.Printf("user-repository.Find(): user not unique! found %v users\n", len(users))
		return nil, errors.New(fmt.Sprintf("user-repository.Find(): user not unique! found %v users\n", len(users)))
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	return users[0], nil
}

func (p UserRepositoryImpl) Create(userId string, user *domain.User) error {
	_, err := p.Find(userId)
	if err != nil {
		_, err := p.connector.Execute(createUserQuery, userId, user.AvatarUrl, user.BulkMode)
		if err != nil {
			log.Printf("user-repository.Create(): error while executing user insert statement: %v\n", err)
			return err
		}
		return nil
	} else {
		return errors.New("user already exists")
	}
}

func (p UserRepositoryImpl) Update(userId string, user *domain.User) error {
	_, err := p.Find(userId)
	if err == nil {
		_, err := p.connector.Execute(updateUserQuery, userId, user.AvatarUrl, user.BulkMode)
		if err != nil {
			log.Printf("user-repository.Update(): error while executing user update statement: %v\n", err)
			return err
		}
		return nil
	} else {
		return errors.New("user does not exists")
	}
}

func (p UserRepositoryImpl) Delete(userId string) error {
	_, err := p.Find(userId)
	if err == nil {
		_, err := p.connector.Execute(deleteUserQuery, userId)
		if err != nil {
			log.Printf("user-repository.Delete(): error while executing user delete statement: %v\n", err)
			return err
		}
		return nil
	} else {
		return errors.New("user does not exists")
	}
}
