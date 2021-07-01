package Models

import (
	"fmt"

	"api-gin-gionic-gorm-mysql/Configuration"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {

	if err = Configuration.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Configuration.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Configuration.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Configuration.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Configuration.DB.Where("id = ?", id).Delete(user)
	return nil
}
