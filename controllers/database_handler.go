package controllers

import (
	"log"

	"github.com/ValonRexhepi/Book-Management-System-REST/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect method that will initialize the connection
// to the database.
func Connect() {
	connection, err := gorm.Open(mysql.Open("root:secret@/BookDB"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Connection to database failed: ", err)
	}
	DB = connection
}

// Migrate method that will migrate the schemas
// to the database.
func Migrate() {
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalln("Schema migration failed: ", err)
	}
}
