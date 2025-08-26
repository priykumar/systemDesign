package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/service"
)

type SubscriberController struct {
	svc service.SubscriberService
}

func NewSubController(svc service.SubscriberService) *SubscriberController {
	return &SubscriberController{svc}
}

func (s *SubscriberController) NewSubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sName := vars["sub"]

	err := s.svc.AddSubscriber(sName)
	txt := ""
	if err != nil {
		txt = err.Error()
	}

	json.NewEncoder(w).Encode(txt)
}
