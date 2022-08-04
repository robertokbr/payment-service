package repositories

import "github.com/robertokbr/payment-service/src/domain/entities"

type UsersRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
