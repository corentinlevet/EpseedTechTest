package db

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Salt      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

const saltSize = 32

func GenerateSalt() (string, error) {
	saltBytes := make([]byte, saltSize)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(saltBytes), nil
}

func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return hashedPassword
}

func CreateUser(username, email, password string) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword := HashPassword(password, salt)
	user := User{Username: username, Email: email, Password: hashedPassword, Salt: salt}
	result := DbInstance.Create(&user)
	return result.Error
}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			hashedPassword := HashPassword(password, user.Salt)
			if hashedPassword == user.Password {
				return &user, nil
			}
		}
	}

	return nil, nil
}

func GetUserByEmail(email string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}

func GetUsers() ([]User, error) {
	var users []User
	result := DbInstance.Find(&users)
	return users, result.Error
}

func UpdateUser(id int, newUsername, newEmail string) error {
	user := User{Username: newUsername, Email: newEmail}
	result := DbInstance.Model(&User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func DeleteUser(id int) error {
	result := DbInstance.Delete(&User{}, id)
	return result.Error
}
