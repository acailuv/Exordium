package user

import "main/database/models"

func (o *userOrmer) SelectAll() ([]models.User, error) {

	users := make([]models.User, 0)

	err := o.db.Model(&users).Select()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}
