package models

import (
	"github.com/jinzhu/gorm"
	"go_casbin/pkg/logger"
)

func GetAllUser() ([]*User, error) {
	var user []*User
	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(data map[string]interface{}) (id uint, err error) {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func UpdateUser(user UserRegister) ( err error) {

	if err := db.Model(&user).Update(UserRegister{user.Username, user.Password, user.CompanyID, user.NickName}).Error; err != nil {
		return err
	}
	return nil
}

//通过用户和密码查询用户
func CheckUser(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(&User{Username: username, Password: password}).First(&user).Error
	logger.Debug("CheckUser", user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func DeleteUser(user UserRegister) error {
	//	///先查用戶是否存在
	//存在，然后在数据库删除
	//var user User
	err := db.Delete(&user).Error
	return err
}

func GetOneUser(username string) (*User,error) {
	var user *User
	err:= db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
