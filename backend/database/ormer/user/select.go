package user

import (
	"main/database/models"

	"github.com/go-pg/pg/v10"
)

func (o *userOrmer) Select(ID string) (models.User, error) {

	user := models.User{}

	err := o.db.Model(&user).
		Where("id = ?", ID).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}
