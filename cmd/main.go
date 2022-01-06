package main

import (
	"calendar/internal/server"
	"calendar/internal/server/v1/handler"
	"calendar/internal/storage/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	_   = godotenv.Load()
	err error
)

func main() {
	s := server.NewServer()
	handler.Database, err = db.NewDb(os.Getenv("DSN"))
	if err != nil {
		log.Println(err)
	}
	err := s.Run()
	if err != nil {
		logrus.Fatalf("An error has occurred: %s", err.Error())
		return
	}

}
