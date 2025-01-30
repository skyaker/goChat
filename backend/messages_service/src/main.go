package main

import (
	"database/sql"
	"fmt"
	"messages_service/handlers"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// ----------------------------------------- Dev message test --------------------------------------------------

	newmes := handlers.NewMessage{
		DialogId:  1,
		SenderId:  1,
		Content:   "Hello Andrew",
		CreatedAt: time.Now(),
	}
	newmes.SendMessage(db)

	// ----------------------------------------- Dev message test --------------------------------------------------

	delmes := handlers.DeletedMessage{
		MessageId: 21,
	}
	delmes.DeleteMessage(db)

	// ----------------------------------------- Dev message test --------------------------------------------------

	newmes = handlers.NewMessage{
		DialogId:  1,
		SenderId:  1,
		Content:   "Hello Andrew",
		CreatedAt: time.Now(),
	}
	newmes.SendMessage(db)

	editmes := handlers.EditedMessage{
		MessageId:  23,
		NewContent: "Hello Alex",
	}
	editmes.EditMessage(db)
}
