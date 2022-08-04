package main

import (
	"flag"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/robertokbr/payment-service/src/application/servers"
	"github.com/robertokbr/payment-service/src/domain/usecases"
	"github.com/robertokbr/payment-service/src/infra/database/postgres/repositories"
)

var db *gorm.DB

func main() {
	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()
	log.Printf("Choosen port: %v", port)

}

func setupUserServer() *servers.UsersServer {
	usersRepository := repositories.UsersRepositoryPG{db}
	usersUsecase := usecases.NewUsersUsecase(&usersRepository)
	usersServer := servers.NewUsersServer(usersUsecase)
}
