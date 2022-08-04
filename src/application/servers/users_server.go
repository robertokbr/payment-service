package servers

import (
	"context"
	"log"

	"github.com/robertokbr/payment-service/src/domain/usecases"
	"github.com/robertokbr/payment-service/src/infra/pb"
)

type UsersServer struct {
	usersUsecase usecases.UsersUsecase
}

func NewUsersServer(usersUsecase *usecases.UsersUsecase) *UsersServer {
	return &UsersServer{usersUsecase}
}

func (u *UsersServer) Create(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	user, err := u.usersUsecase.Create(name, email, password)

	if err != nil {
		log.Fatalf("[GRPC_SERVER] Error while creating a user: %v", err)
	}

	response := pb.CreateUserResponse{
		Token: user.Token,
	}

	return &response, nil
}
