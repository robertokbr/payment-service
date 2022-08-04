package entities

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key;DEFAULT uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deletedAt" gorm:"type:timestamp" sql:"index"`
}

// Function called by gorm
func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during creating entity: %v", err)
		return err
	}

	err = scope.SetColumn("UpdatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during creating entity: %v", err)
		return err
	}

	err = scope.SetColumn("Id", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during creating entity: %v", err)
		return err
	}

	return nil
}
