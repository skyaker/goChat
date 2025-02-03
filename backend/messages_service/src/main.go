package main

import (
	"database/sql"
	"fmt"
	"log"
	handlers "messages_service/src/handlers"
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

	newmes := handlers.MessageCreate{
		DialogId:  1,
		SenderId:  1,
		Content:   "Hello Andrew",
		CreatedAt: time.Now(),
	}
	err = newmes.SendMessage(db)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	// ----------------------------------------- Dev message test --------------------------------------------------

	delmes := handlers.MessageDelete{
		MessageId: 21,
	}
	err = delmes.DeleteMessage(db)
	if err != nil {
		log.Printf("Failed to delete message: %v", err)
	}

	// ----------------------------------------- Dev message test --------------------------------------------------

	newmes = handlers.MessageCreate{
		DialogId:  1,
		SenderId:  1,
		Content:   "Hello Andrew",
		CreatedAt: time.Now(),
	}
	newmes.SendMessage(db)

	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	editmes := handlers.MessageEdit{
		MessageId:  23,
		NewContent: "Hello Alex",
	}
	editmes.EditMessage(db)

	if err != nil {
		log.Printf("Failed to edit message: %v", err)
	}
}
