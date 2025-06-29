package main

import (
	"golang-solid-clean-architecture/config"
	"golang-solid-clean-architecture/database"
	"golang-solid-clean-architecture/infrastructures"
	httpHandler "golang-solid-clean-architecture/interfaces/http"
	"golang-solid-clean-architecture/usecases"
	"log"
	"net/http"
)

func methodMiddleware(handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			http.Error(writer, "Method now allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(writer, request)
	}
}

func main() {
	dbConfig := config.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		DBName:   "golang_learnings",
	}

	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := infrastructures.NewUserRepository(db)
	usecase := usecases.NewUserUsecase(repo)
	handler := httpHandler.NewUserHandler(usecase)

	http.HandleFunc("/register", methodMiddleware(handler.Register, http.MethodPost))
	http.HandleFunc("/user", methodMiddleware(handler.GetUser, http.MethodGet))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
