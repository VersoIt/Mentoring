package main

import (
	"fmt"
	localhttp "testing/internal/delivery/http"
	"testing/internal/repository"
	"testing/internal/service"
)

func main() {
	personRepo := repository.NewPersonRepository()
	personService := service.NewPersonService(personRepo)
	handler := localhttp.NewHandler(personService)

	handler.InitRoutes()
	err := handler.Run()
	if err != nil {
		fmt.Println(err)
	}
}
