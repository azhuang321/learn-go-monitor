package models

import "github.com/jinzhu/gorm"

// User 用户表 model 定义
type User struct {
	gorm.Model
	UserName string `gorm:"column:name;type:varchar(100);unique_index;default:null"`
	Password string `gorm:"column:password;type:varchar(100);default:null"`
	Email    string `gorm:"column:email;type:varchar(100);default:null"`
	Avatar   string `gorm:"column:avatar;type:varchar(100);default:null"`
	Status   string `gorm:"column:status;type:ENUM('ENABLE','DISABLE')"`
}

// Insert 新增用户
func (u *User) Insert() (userId uint, err error) {
	result := DB.Create(&u)
	userId = u.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询用户详情
func (u *User) FindOne(condition map[string]interface{}) (*User, error) {
	var userInfo User
	result := DB.Select("id,name,email,avatar,password").Where(condition).First(&userInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if userInfo.ID > 0 {
		return &userInfo, nil
	}
	return nil, nil
}

// FindAll 获取用户列表
func (u *User) FindAll(pageNum int, pageSize int, condition interface{}) (users []User, err error) {
	result := DB.Offset(pageNum).Limit(pageSize).Select("id", "name", "email").Where(condition).Find(&users)
	err = result.Error
	return
}

// UpdateOne 修改用户
func (u *User) UpdateOne(userId uint, data map[string]interface{}) (*User, error) {
	err := DB.Model(&User{}).Where("id = ?", userId).Update(data).Error
	if err != nil {
		return nil, err
	}
	var updUser User
	err = DB.Select([]string{"id", "name", "email", "avatar"}).First(&updUser, userId).Error
	if err != nil {
		return nil, err
	}
	return &updUser, nil
}

// DeleteOne 删除用户
func (u *User) DeleteOne(userID uint) (delUser User, err error) {
	if err = DB.Select([]string{"id"}).First(&u, userID).Error; err != nil {
		return
	}

	if err = DB.Delete(&u).Error; err != nil {
		return
	}
	delUser = *u
	return
}
