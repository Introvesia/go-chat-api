package libs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Define connection URL
	connString := "postgres://" + GetEnv("POSTGRES_USER") + ":" +
		GetEnv("POSTGRES_PASSWORD") + "@" + GetEnv("POSTGRES_HOST") + ":" +
		GetEnv("POSTGRES_PORT") + "/" + GetEnv("POSTGRES_DB")

	conn, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	return conn
}
