package services

import (
	"github.com/cholid97/go-kredit/dto"
	"github.com/cholid97/go-kredit/repositories"

	"github.com/cholid97/go-kredit/models"
)

type UserService interface {
	GetAllCredits() ([]dto.UserContractResponse, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	CreateContract(contract *models.Contract) (*models.Contract, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllCredits() ([]dto.UserContractResponse, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) CreateContract(contract *models.Contract) (*models.Contract, error) {
	ctr, err := s.repo.CreateCredit(contract)

	return ctr, err
}
