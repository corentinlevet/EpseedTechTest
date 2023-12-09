package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier .env:", err)
		return
	}
}

func InitTimezone() {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		fmt.Println("Erreur lors du chargement du timezone:", err)
		return
	}

	time.Local = loc
}
