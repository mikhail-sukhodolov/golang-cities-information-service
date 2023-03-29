package handler

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var s model.CityRequest
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: New model %v", string(content)))

	err = json.Unmarshal(content, &s)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(&s)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("The model of %v was created with the id %v", s.Name, id)
	newMessageResponse(w, http.StatusCreated, message)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Delete(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("City with identifier %v deleted", id)
	newMessageResponse(w, http.StatusOK, message)
}

func (h *Handler) GetFull(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Full info %v", id))

	city, err := h.services.GetFull(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message, err := json.Marshal(city)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("Sending: %v", string(message)))
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}

func (h *Handler) SetPopulation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("PUT: Update population model %v", id))

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var population model.SetPopulationRequest
	err = json.Unmarshal(content, &population)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.SetPopulation(id, population.Population)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newMessageResponse(w, http.StatusOK, "Change was successful")
}

func (h *Handler) GetFromRegion(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")

	log.Info().Msg(fmt.Sprintf("GET: Сities by region %v", region))

	cityNames, err := h.services.GetFromRegion(region)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}

func (h *Handler) GetFromDistrict(w http.ResponseWriter, r *http.Request) {
	district := chi.URLParam(r, "district")

	log.Info().Msg(fmt.Sprintf("GET: Сities by district %v", district))

	cityNames, err := h.services.GetFromDistrict(district)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}

func (h *Handler) GetFromPopulation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	fmt.Println("test", content)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Сities by population range %v", string(content)))

	var populationRange model.RangeRequest
	err = json.Unmarshal(content, &populationRange)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.services.GetFromPopulation(populationRange.Start, populationRange.End)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}

func (h *Handler) GetFromFoundation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Сities by foundation range %v", string(content)))

	var foundationRange model.RangeRequest
	err = json.Unmarshal(content, &foundationRange)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.services.GetFromFoundation(foundationRange.Start, foundationRange.End)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}
