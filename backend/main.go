package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// "backend/database"
)

func initDB() {
	dsn := "root:python101@tcp(127.0.0.1:3306)/ContentManager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	log.Println(db)
	log.Println("DB Connected!")
}

func main() {
	app := fiber.New()

	initDB()

	// database.DBConn.AutoMigrate(&book.Book{})
	// log.Println("DB Migrated")

	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })


	log.Fatal(app.Listen(":8000"))
}

/*
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"tutorial/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
*/