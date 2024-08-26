// apps/server/src/app/db.go
package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

type Synonym struct {
	Phrase   string `gorm:"primaryKey"`
	Synonyms []byte `gorm:"type:jsonb"`
}

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// Run migrations
	if err = DB.AutoMigrate(&Synonym{}); err != nil {
		log.Fatalf("Error running migrations: %q", err)
	}
}
