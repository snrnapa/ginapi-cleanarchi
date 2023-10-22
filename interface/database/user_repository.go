package database

import (
	"ginapi-cleanarchi/domain/entity"
	"ginapi-cleanarchi/domain/repository"

	"github.com/go-xorm/xorm"
)

type UserRepository struct {
	*xorm.Engine
}

func NewUserRepository(db *xorm.Engine) repository.IUserRepository {
	return &UserRepository{db}

}

func (r *UserRepository) SelectAllUsers(userId string) (*entity.User, error) {

	user := entity.User{}
	_, error := r.Table('users')

}
