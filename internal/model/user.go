package model

import (
	"webDemo/pkg/errcode"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32 `json:"id"`
	Name      string `gorm:"primary_key" json:"name"`
	Password  string `json:"password"`
	Privilege uint32 `json:"privilege"`
}

func (user *User) TableName() string {
	return "user"
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", u.ID).Update(u).Error
}

func (u *User) DeleteById(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(&u).Error
}

func (u *User) DeleteByName(db *gorm.DB) error {
	return db.Where("name = ?", u.Name).Delete(&u).Error
}

func (u *User) Count(db *gorm.DB) (int, error) {
	count := 0
	if u.Name != "" {
		db = db.Where("name = ?", u.Name)
	}
	if err := db.Model(&u).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u *User) GetPasswd(db *gorm.DB) (string, error) {
	var user User
	result := db.Where("name = ?", u.Name).First(&user)

	if result.Error != nil {
		return "", errcode.ErrorUserNotExist
	}
	return user.Password, nil
}
