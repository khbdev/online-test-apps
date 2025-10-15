package service

import (
	models "user-service/internal/model"
	repository "user-service/internal/repostory"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}


func (s *UserService) Create(
	firstName, lastName, phone, email, tgUsername,
	bolimlar, savollar, javoblar string,
	togriJavoblar, natogriJavoblar int,
) (*models.User, error) {

	user := &models.User{
		FirstName:       firstName,
		LastName:        lastName,
		Phone:           phone,
		Email:           email,
		TgUsername:      tgUsername,
		Bolimlar:        bolimlar,
		Savollar:        savollar,
		Javoblar:        javoblar,
		TogriJavoblar:   togriJavoblar,
		NatogriJavoblar: natogriJavoblar,
	}

	return s.repo.Create(user)
}


func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}


func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}


func (s *UserService) Update(
	id uint,
	firstName, lastName, phone, email, tgUsername,
	bolimlar, savollar, javoblar string,
	togriJavoblar, natogriJavoblar int,
) (*models.User, error) {

	updatedUser := &models.User{
		FirstName:       firstName,
		LastName:        lastName,
		Phone:           phone,
		Email:           email,
		TgUsername:      tgUsername,
		Bolimlar:        bolimlar,
		Savollar:        savollar,
		Javoblar:        javoblar,
		TogriJavoblar:   togriJavoblar,
		NatogriJavoblar: natogriJavoblar,
	}

	return s.repo.Update(id, updatedUser)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
