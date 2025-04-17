package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Run() error {
	port := 7540

	// Проверяем переменную окружения TODO_PORT
	if portEnv := os.Getenv("TODO_PORT"); portEnv != "" {
		if parsedPort, err := strconv.Atoi(portEnv); err == nil {
			port = parsedPort
		} else {
			log.Printf("Недопустимый порт: %s, использование порта по умолчанию: %d", portEnv, port)
		}
	}

	// Логируем, на каком порту запущен сервер
	log.Printf("Сервер запущен на порту: %d", port)

	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
