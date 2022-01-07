package user

import (
	"main/database/models"
)

func (o *userOrmer) Update(user models.User) error {
	_, err := o.db.Model(&user).WherePK().UpdateNotZero()
	if err != nil {
		return err
	}

	return nil
}
