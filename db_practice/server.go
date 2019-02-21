package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Aakash struct {
	gorm.Model
	Weight uint
	Beard  string
}

func main() {
	db, err := gorm.Open("sqlite3", "db_practice.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Aakash{})

	// Create
	db.Create(&Aakash{Weight: 170, Beard: "Awesome"})

	// Read
	var aakash Aakash
	db.First(&aakash, 1)
	db.First(&aakash, "beard = ?", "Awesome")

	var aakashes []Aakash
	db.Find(&aakashes)
	fmt.Println("All Aakashes are: ", aakashes)
	// // Update
	db.Model(&aakash).Update("Beard", "Super Awesome")

	// // Delete
	db.Delete(&aakash)
}
