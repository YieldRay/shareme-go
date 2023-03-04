package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"shareme/db"
	. "shareme/public"
	"shareme/routes"
)

func isSet(str ...string) bool {
	for _, value := range str {
		if len(value) == 0 {
			return false
		}
	}
	return true
}

func getEnv(key string, defaults ...string) string {
	e := os.Getenv(key)
	if len(e) == 0 {
		if len(defaults) == 0 {
			return ""
		}
		return defaults[0]
	}
	return e
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.SetTrustedProxies([]string{"192.168.1.1"})

	// attach static
	efs := Public
	// config database
	godotenv.Load()
	MONGODB_URI := getEnv("MONGODB_URI")
	MONGODB_NAME := getEnv("MONGODB_NAME")
	MONGODB_COLLECTION := getEnv("MONGODB_COLLECTION")
	MYSQL_USERNAME := getEnv("MYSQL_USERNAME")
	MYSQL_PASSWORD := getEnv("MYSQL_PASSWORD")
	MYSQL_HOST := getEnv("MYSQL_HOST", "127.0.0.1")
	MYSQL_PORT := getEnv("MYSQL_PORT", "3306")
	MYSQL_DB_NAME := getEnv("MYSQL_DB_NAME")
	MYSQL_TABLE_NAME := getEnv("MYSQL_TABLE_NAME")

	var database db.IDB
	if isSet(MONGODB_URI, MONGODB_NAME, MONGODB_COLLECTION) {
		fmt.Println("Using MongoDB")
		database = db.MongoDB(MONGODB_URI, MONGODB_NAME, MONGODB_COLLECTION)
	} else if isSet(MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB_NAME, MYSQL_TABLE_NAME) {
		fmt.Println("Using MySQL")
		database = db.MySQL(MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB_NAME, MYSQL_TABLE_NAME)
	} else {
		fmt.Println("Using TMP Cache")
		database = db.TmpDB()
	}

	// attach middleware
	routes.APIMiddleware(app, database)
	routes.StaticMiddleware(app, database, efs)
	/// config port
	port := getEnv("PORT", "8080")
	fmt.Println("Server listen at http://localhost:" + port)
	app.Run(":" + port)
}
