package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string `json:"-"`
}

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

func DeleteUser(db *gorm.DB, id int) (*User, error) {
	var u User
	if err := db.First(&u, id).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
