package database

import "CRUD/model"

func GetUsers(users *[]model.Contact) error {
	tx := connector.Find(&users)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetUser(user *model.Contact, params map[string]string) error {
	tx := connector.First(&user, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func CreateUser(user *model.Contact) error {
	tx := connector.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateUser(user *model.Contact, params map[string]string) error {
	tx := connector.First(&user, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteUser(user *model.Contact, params map[string]string) error {
	tx := connector.Delete(&user, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
