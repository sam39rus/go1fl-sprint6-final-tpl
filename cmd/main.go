package main

import (
	"log"

	"go1fl-sprint6-final-tpl/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(log.Writer(), "[SERVER] ", log.LstdFlags|log.Lshortfile)

	// Создаем сервер с указанным логгером
	srv := server.CreateServer(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}
