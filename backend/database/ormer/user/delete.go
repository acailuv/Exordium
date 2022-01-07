package user

import "main/database/models"

func (o *userOrmer) Delete(ID string) error {

	user := models.User{}

	_, err := o.db.Model(&user).
		Where("id = ?", ID).
		Delete()
	if err != nil {
		return err
	}

	return nil
}
