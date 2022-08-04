package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/robertokbr/payment-service/src/domain/entities"
)

type UsersRepositoryPG struct {
	Db *gorm.DB
}

func (u *UsersRepositoryPG) FindByEmail(email string) (*entities.User, error) {
	return entities.NewUser(), nil
}

func (r *UsersRepositoryPG) Create(
	user *entities.User,
) (*entities.User, error) {
	err := r.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error during user creation: %v", err)
		return user, err
	}

	return user, nil
}
