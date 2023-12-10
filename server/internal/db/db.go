package db

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

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

func ConnectToMariaDB() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := dbUser + ":" + dbPassword + "@tcp(database:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Connecté à MariaDB!")
	return db, nil
}
