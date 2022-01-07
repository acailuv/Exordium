package user

import (
	"main/database/models"
)

func (o *userOrmer) Insert(user models.User) error {
	_, err := o.db.Model(&user).Insert()
	if err != nil {
		return err
	}

	return nil
}
