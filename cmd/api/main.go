package main

import (
	"net/http"

	"github.com/johnfercher/go-hexagonal/internal/adapters/drivens/memdb"
	"github.com/johnfercher/go-hexagonal/internal/adapters/drivers/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnfercher/go-hexagonal/internal/services"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	repository := memdb.NewUserRepository()
	service := services.NewUserRegister(repository)
	handler := rest.NewUserRegisterHandler(service)

	r.MethodFunc(handler.Method(), handler.Pattern(), handler.Func)

	http.ListenAndServe(":3001", r)
}
