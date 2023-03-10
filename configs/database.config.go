package configs

import (
	"fmt"
	"kanban-app/database"
	"kanban-app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_name = "railway"
var db_port = "6483"
var db_user = "root"
var db_password = "kVoTkFv55d5Sfk7Dj8Pq"
var db_host = "containers-us-west-32.railway.app"

func BootDatabase() {
	if dbPortEnv := os.Getenv("MYSQLPORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbNameEnv := os.Getenv("MYSQLDATABASE"); dbNameEnv != "" {
		db_name = dbNameEnv
	}

	if dbUserEnv := os.Getenv("MYSQLUSER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("MYSQLPASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("MYSQLHOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func LoadConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverPort := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbName := os.Getenv("MYSQLDATABASE")

	config := models.Config{
		ServerPort: serverPort,
		Database: models.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
		},
	}
	return config
}

func ConnectDatabase() {

	var errConnection error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	database.DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Database connected")
	}

}

func RunMigration() {
	err := database.DB.AutoMigrate(
		models.Todo{},
	)

	if err != nil {
		fmt.Println(err)
		log.Println("Failed to migrate schema")
	} else {
		log.Println("schemas migrated")
	}
}
