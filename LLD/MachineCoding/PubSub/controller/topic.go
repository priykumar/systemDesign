package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/service"
)

type TypeController struct {
	svc service.TopicService
}

func NewTypeController(svc service.TopicService) *TypeController {
	return &TypeController{svc}
}

func (t *TypeController) CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topicName := vars["topic"]

	err := t.svc.CreateTopic(topicName)
	txt := ""
	if err != nil {
		txt = err.Error()
	}
	json.NewEncoder(w).Encode(txt)
}

func (t *TypeController) AddSubsHandler(w http.ResponseWriter, r *http.Request) {
	type s struct {
		Name string `json:"name"`
		Sub  string `json:"sub"`
	}

	var req s
	json.NewDecoder(r.Body).Decode(&req)
	err := t.svc.AddSubsToTopic(req.Name, req.Sub)
	txt := ""
	if err != nil {
		txt = err.Error()
	}
	json.NewEncoder(w).Encode(txt)
}
