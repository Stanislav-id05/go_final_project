package main

import (
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"

	"github.com/Stanislav-id05/go_final_project/pkg/api"
)

func main() {
	// Запуск сервера иниц.
	api.Init()
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()
	server.Run()
}
