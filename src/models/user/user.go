package user

type User struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string `json:"-"`
}
