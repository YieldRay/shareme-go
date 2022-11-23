package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"shareme/db"
	"shareme/route"
)

func isSet(str ...string) bool {
	for _, value := range str {
		if len(value) == 0 {
			return false
		}
	}
	return true
}

//go:embed public/*
var efs embed.FS

func main() {
	gin := gin.New()
	gin.SetTrustedProxies([]string{"192.168.1.1"})
	godotenv.Load()
	MONGO_DB_URI := os.Getenv("MONGO_DB_URI")
	MONGO_DB_NAME := os.Getenv("MONGO_DB_NAME")
	MONGO_DB_COLLECTION := os.Getenv("MONGO_DB_COLLECTION")
	var database db.IDB
	if isSet(MONGO_DB_URI, MONGO_DB_NAME, MONGO_DB_COLLECTION) {
		fmt.Println("Using MongoDB")
		database = db.MongoDB(MONGO_DB_URI, MONGO_DB_NAME, MONGO_DB_COLLECTION)
	} else {
		fmt.Println("Using TMP Cache")
		database = db.TmpDB()
	}

	route.APIMiddleware(gin, database)
	route.StaticMiddleware(gin, database, efs)
	gin.Run()
}
