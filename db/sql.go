package db

import (
	"log"
	"os"
	"path"

	"github.com/bin3xish477/csg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	hd, _  = os.UserHomeDir()
	Folder = path.Join(hd, ".csg")
	Vault  = path.Join(Folder, "vault.db")

	DB *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open(Vault), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}

	DB = db
}

func CreateDB() {
	f, err := os.Create(Vault)
	if err != nil {
		log.Printf("unable to create file in folder: %s", Folder)
	}
	f.Close()

	err = os.Chmod(Vault, 0755)
	if err != nil {
		log.Println("couldn't change database permissions...")
	}

	Connect()

	DB.AutoMigrate(&models.Credential{})
}
