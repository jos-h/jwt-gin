package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jwt-gin/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}


func (user *User) SaveUser() (*User, error) {
	err := DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, err
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))
	return nil
}

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(user_name string, password string) (string, error) {
	var err error
	user := User{}
	if err = DB.Take(&user, "user_name = ?", user_name).Error; err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserByID(user_id uint) (User, error) {

	var u User

	if err := DB.First(&u, user_id).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func GetUsers() ([]User, error) {
	users := []User{}

	if err := DB.Order("created_at desc").Find(&users).Error; err != nil {
		return users, err
	}
	for i := range users{
		users[i].PrepareGive()
	}
	return users, nil
}
