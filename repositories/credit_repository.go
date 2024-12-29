package repositories

import (
	"errors"
	"time"

	"github.com/cholid97/go-kredit/dto"
	"github.com/cholid97/go-kredit/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]dto.UserContractResponse, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) error
	CreateCredit(contract *models.Contract) (*models.Contract, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]dto.UserContractResponse, error) {
	var users []dto.UserContractResponse
	err := r.db.Table("users AS a").
		Select("a.full_name, b.asset, b.installment_amount, b.installment_interest").
		Joins("INNER JOIN contracts AS b ON a.id = b.user_id").
		Limit(100).Scan(&users).Error

	return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) CreateCredit(contract *models.Contract) (*models.Contract, error) {
	var limit models.Limit

	tx := r.db.Begin()
	err := tx.Where("user_id = ? AND MONTH(month_limit) = ? ", contract.UserId, int(time.Now().Month())).Find(&limit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}

		return nil, err
	}

	if limit.MonthlyLimit < contract.OTR {
		return nil, errors.New("Limit exceeded")
	}

	err = tx.Create(contract).Error

	if err != nil {
		return nil, err
	}

	limit.MonthlyLimit -= contract.OTR

	err = tx.Save(&limit).Error

	if err != nil {
		tx.Rollback()
		return nil, err

	}

	tx.Commit()

	return contract, err

}
