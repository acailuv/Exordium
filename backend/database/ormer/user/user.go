package user

import (
	"main/database/models"

	"github.com/go-pg/pg/v10"
)

type UserOrmer interface {
	Select(ID string) (models.User, error)
	SelectAll() ([]models.User, error)
	Insert(user models.User) error
	Update(user models.User) error
	Delete(ID string) error
}

type userOrmer struct {
	db *pg.DB
}

func NewUserOrmer(db *pg.DB) UserOrmer {
	return &userOrmer{
		db: db,
	}
}
