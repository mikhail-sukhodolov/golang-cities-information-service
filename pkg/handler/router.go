package handler

import (
	"github.com/go-chi/chi"
	"golang-cities-information-service/pkg/service"
)

type Handler struct {
	services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/create", h.Create)
	router.Delete("/{id}", h.Delete)
	router.Get("/{id}", h.GetFull)
	router.Put("/population/{id}", h.SetPopulation)
	router.Get("/region/{region}", h.GetFromRegion)
	router.Get("/district/{district}", h.GetFromDistrict)
	router.Get("/population/range", h.GetFromPopulation)
	router.Get("/foundation/range", h.GetFromFoundation)

	return router
}
