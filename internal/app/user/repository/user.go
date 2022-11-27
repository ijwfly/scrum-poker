package repository

import (
	"scrum-poker/internal/postgres"
)

type UserRepository struct {
	*postgres.PostgresDB
}

//func (p *UserRepository) GetUser(userId int64) (user.User.User, error) {
//}
