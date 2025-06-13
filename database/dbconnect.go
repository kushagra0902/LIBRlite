package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	
)

var Pool *pgxpool.Pool // we want to use Pool in another file. Here we just want to declare. Therefore is didnt declare first and then assign, it will throw error that it is declared and nt used

func ConnectDB(databaseURL string) {
	var err2 error
	Pool, err2 = pgxpool.New(context.Background(), databaseURL)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err2)
		os.Exit(1)
	}
	
	fmt.Println("Connected to Database Successfully")

}
