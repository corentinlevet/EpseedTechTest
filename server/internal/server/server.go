package server

import (
	"epseed/internal/db"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func InitServer() {
	var nbTry int = 0
	var err error
	for nbTry < 3 {
		fmt.Println("Tentative {", (nbTry + 1), "} de connexion à la base de donnée...")
		db.DbInstance, err = db.ConnectToMariaDB()
		if err != nil {
			fmt.Println("Tentative de connexion à la base de donnée échouée, nouvelle tentative dans 5 secondes...")
			nbTry++
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	if nbTry == 3 {
		fmt.Println("Erreur de connexion à la base de données après 3 tentatives, arrêt du serveur...")
		return
	}
	sqlDB, err := db.DbInstance.DB()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données:", err)
		return
	}
	defer sqlDB.Close()

	InitRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
