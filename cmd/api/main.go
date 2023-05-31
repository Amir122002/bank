package main

import (
	"flag"
	"github.com/Amir122002/bank/pkg/handlers"
	"github.com/Amir122002/bank/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	DBName := flag.String("dbname", "bank", "Enter the name of DB")
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

	router.GET("/Get", h.GetAllUser)
	router.GET("/users/:login", h.GetUserByLogin)

	router.GET("/users/:login/replenish-money/:replenish", h.ReplenishUserMoney)
	router.GET("/users/:login/withdraw-money/:withdraw", h.WithdrawUserMoney)

	router.Run(":4000")
}

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Cell{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
