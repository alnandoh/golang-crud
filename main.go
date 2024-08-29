package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"
)

func main() {
	fmt.Printf("Start Application")
	//database
	db := config.DatabaseConnection()

	//repository
	bookRepository := repository.NewBookRepository(db)

	//service
	bookService := service.NewBookServiceImpl(bookRepository)

	//controller
	bookController := controller.NewBookController(bookService)

	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8080", Handler: routes}

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}
