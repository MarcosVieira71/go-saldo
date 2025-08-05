package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	return u, nil
}

func AddUser(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var u User
	err := db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByID(db *gorm.DB, id int) (*User, error) {
	var u User
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
