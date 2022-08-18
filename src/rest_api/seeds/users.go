package seeds

import (
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, id int, firstname string, lastname string, email string) error {
	return db.Create(&model.User{ID: id, Firstname: firstname, Lastname: lastname, Email: email}).Error
}
