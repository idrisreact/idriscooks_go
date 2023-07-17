package database

import (
	"fmt"
	"idrisgo/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DBInstanse struct {
	DB *gorm.DB
}

var Db DBInstanse

func ConnectDb(){

	dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
)
	db ,err :=gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil{
		log.Fatal("failed to connect to database. \n",err)
		os.Exit(2)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Recipe{})

	Db = DBInstanse{
		DB: db,
	}
}	