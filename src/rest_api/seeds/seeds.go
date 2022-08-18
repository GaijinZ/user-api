package seeds

import (
	"github.com/GaijinZ/user-api/src/rest_api/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, 1, "Jane", "Milo", "jane.milo@gmail.com")

			},
		},

		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, 2, "John", "Milo", "john.milo1@gmail.com")
			},
		},
	}
}
