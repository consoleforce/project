// services/contact/main.go

package main

import (
	"github.com/consoleforce/project/pkg/store/postgres" // Подключаем пакет с реализацией PostgreSQL
	"github.com/consoleforce/project/services/contact/internal"
	"github.com/consoleforce/project/services/contact/internal/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация подключения к БД PostgreSQL
	db, err := postgres.ConnectDB("localhost", "5432", "user", "password", "dbname")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	contactRepo := repository.NewContactRepository(db)
	groupRepo := repository.NewGroupRepository(db)

	// Создание конструктора для слоев
	constructor := internal.NewConstructor(contactRepo, groupRepo)

	// Создание роутера для обработки HTTP запросов
	r := mux.NewRouter()

	// Инициализация и привязка обработчиков к роутеру
	constructor.ContactDel.BindHandlers(r)

	// Запуск HTTP сервера
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
