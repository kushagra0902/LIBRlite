package main

import (
	"github/Kushagra0902/LIBRlite/database" // suppose in you go.mod module abc is used. Then to import files you have to do abc/filename
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load()
	if err!=nil{
		panic(err)
	}
	databaseURL := os.Getenv("DATABASE_URL")
	database.ConnectDB(databaseURL)
}
