package models

import (
	"epseed/internal/db"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Salt      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func GetUsers() ([]*User, error) {
	var users []*User
	err := db.DbInstance.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DbInstance.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			hashedPassword := db.HashPassword(password, user.Salt)
			if hashedPassword == user.Password {
				return user, nil
			}
		}
	}

	return nil, nil
}

func CreateUser(username, email, password string) error {
	salt, err := db.GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword := db.HashPassword(password, salt)
	user := User{Username: username, Email: email, Password: hashedPassword, Salt: salt}
	result := db.DbInstance.Create(&user)
	return result.Error
}
