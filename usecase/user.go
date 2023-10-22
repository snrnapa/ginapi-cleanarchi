package usecase

import (
	"ginapi-cleanarchi/common/dto"
	"ginapi-cleanarchi/domain/repository"
)

type UserInputPort interface {
	GetUser(userId string) (*dto.User, error)
}

type UserCreate struct {
	repository.IUserRepository
}

func NewUserCreate(repo repository.IUserRepository) UserInputPort {
	return &UserCreate{repo}

}

func (s *UserCreate) GetUser(userid string) (*dto.User, error) {
	userE, err := s.IUserRepository.SelectAllUsers(userid)
	//エラーが発生している場合は、エラーを返す
	if err != nil {
		return nil, err
	}

	userDto := dto.NewUserByJson(userE.Id, userE.FirstName, userE.LastName)

	return userDto, nil

}
