package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/HarshitPG/ecommerce_api_go/cmd/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	sqlDB, err := sql.Open("mysql", config.Envs.DBUser+":"+config.Envs.DBPassword+"@tcp("+config.Envs.DBAddress+")/"+config.Envs.DBName+"?parseTime=true")
	if err != nil {
		log.Fatalf("could not connect to the MySQL database: %v", err)
	}
	defer sqlDB.Close()

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not create the MySQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("could not create the migrate instance: %v", err)
	}


	if len(os.Args) < 2 {
		log.Fatal("missing migration command (up or down)")
	}
	cmd := os.Args[1]

	
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not run migrate up: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not run migrate down: %v", err)
		}
	default:
		log.Fatalf("unknown command: %s", cmd)
	}
}
