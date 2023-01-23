package models

import (
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string, dbName string, pgUser string) {
	// connect to the postgres db just to be able to run the create db statement
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// if the connection fails because the database does not exist, create it
		if strings.Contains(err.Error(), "database \""+dbName+"\" does not exist") {
			// initiate a temporary connection to the postgres database
			database, err = gorm.Open(postgres.Open(dsn+" dbname=postgres"), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}
			// create the database
			database.Exec("CREATE DATABASE \"" + dbName + "\"")
			database.Exec("GRANT ALL PRIVILEGES ON DATABASE \"" + dbName + "\" TO " + pgUser)
			database.Exec("ALTER DATABASE \"" + dbName + "\" OWNER TO " + pgUser)
			// close the temporary connection
			sql, err := database.DB()
			defer func() {
				_ = sql.Close()
			}()
			if err != nil {
				log.Fatal(err)
			}
			// reconnect to the newly created database
			database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// otherwise we gentlypanic
			panic(err)
		}
	}

	database.AutoMigrate(&Page{})

	DB = database
}

// A function used to properly close the database connection
func Shutdown() error {
	db, _ := DB.DB()
	err := db.Close()
	if err != nil {
		log.Fatal("Failed to close database connection!")
	}

	return err
}
