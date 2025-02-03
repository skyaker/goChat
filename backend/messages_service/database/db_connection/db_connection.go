package connection

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func GetDbConnection() *sql.DB {
	host, _ := os.LookupEnv("host")

	port, _ := os.LookupEnv("port")
	portInt, _ := strconv.Atoi(port)
	portUint := uint(portInt)

	user, _ := os.LookupEnv("user")
	password, _ := os.LookupEnv("password")
	dbname, _ := os.LookupEnv("dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, portUint, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
