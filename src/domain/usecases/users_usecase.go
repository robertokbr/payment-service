package usecases

import (
	"log"

	"github.com/robertokbr/payment-service/src/domain/entities"
	"github.com/robertokbr/payment-service/src/domain/repositories"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase struct {
	usersRepository repositories.UsersRepository
}

func NewUsersUsecase(
	usersRepository repositories.UsersRepository,
) *UsersUsecase {
	return &UsersUsecase{usersRepository}
}

func (u *UsersUsecase) Create(
	name string,
	email string,
	password string,
) (*entities.User, error) {
	foundUser, err := u.usersRepository.FindByEmail(
		email,
	)

	if err != nil {
		log.Fatalf("Error while searching for user by email: %v", err)
		return nil, err
	}

	if foundUser != nil {
		log.Fatal("Email is already in use")
		return nil, err
	}

	user := entities.User{
		Name:  name,
		Email: email,
	}

	bcryptPassword := []byte("")

	bcryptPassword, err = bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return nil, err
	}

	user.Password = string(bcryptPassword)
	user.Token = uuid.NewV4().String()

	return &user, nil
}
