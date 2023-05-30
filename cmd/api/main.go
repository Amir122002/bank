package main

import (
	"flag"
	"log"
)

func main() {
	DBName := flag.String("dbname", "Bank", "Enter the name of DB")
	DBUser := flag.String("dbuser", "postgres", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "17122002amir", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("db connection error:", err)
	}

	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)
	router := gin.Default()

	router.GET("/", h.GetOneUser)

	router.Run(":4000")
}
