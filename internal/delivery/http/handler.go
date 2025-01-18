package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing/internal/model"
)

type PersonService interface {
	Get(name string) (model.Person, error)
}

type Handler struct {
	personService PersonService
}

func NewHandler(personService PersonService) *Handler {
	return &Handler{
		personService: personService,
	}
}

type request struct {
	Name string `json:"name"`
}

func (h *Handler) InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("error while decoding request: %v", err), http.StatusBadRequest)
			return
		}

		person, err := h.personService.Get(req.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("долбоеб ебучий у нас такого нету чела: %s", req.Name), http.StatusNotFound)
			return
		}

		data, err := json.Marshal(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err = w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (h *Handler) Run() error {
	log.Println("starting server on :8080...")
	return http.ListenAndServe(":8080", nil)
}
