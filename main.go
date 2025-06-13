package main

import (
	"fmt"
	"github/Kushagra0902/LIBRlite/database" // suppose in you go.mod module abc is used. Then to import files you have to do abc/filename
	"github/Kushagra0902/LIBRlite/routes"
	"net/http"
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

	r := routes.RouteHandler()
	fmt.Println("Starting Server")
	http.ListenAndServe(":8000", r)
	
}
