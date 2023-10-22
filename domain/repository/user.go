package repository

import "ginapi-cleanarchi/domain/entity"

type IUserRepository interface {
	SelectAllUsers(userId string) (*entity.User, error)
}
