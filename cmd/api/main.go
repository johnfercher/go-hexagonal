package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnfercher/go-hexagonal/internal/adapters/drivens/memdb"
	"github.com/johnfercher/go-hexagonal/internal/adapters/drivers/httphandlers"
	"github.com/johnfercher/go-hexagonal/internal/kycservices"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	repository := memdb.NewUserRepository()
	service := kycservices.NewUserRegister(repository)
	handler := httphandlers.NewUserRegister(service)

	r.MethodFunc(handler.Method(), handler.Pattern(), handler.Func)

	http.ListenAndServe(":3001", r)
}
