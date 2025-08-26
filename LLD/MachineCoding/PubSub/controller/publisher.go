package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/service"
)

type PublisherController struct {
	svc service.PublisherService
}

func NewPubController(svc service.PublisherService) *PublisherController {
	return &PublisherController{svc}
}

func (p *PublisherController) AddPubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pName := vars["pub"]

	err := p.svc.AddPublisher(pName)
	txt := ""
	if err != nil {
		txt = err.Error()
	}

	json.NewEncoder(w).Encode(txt)
}

func (p *PublisherController) PublishHandler(w http.ResponseWriter, r *http.Request) {
	type s1 struct {
		T string `json:"topic"`
		C string `json:"content"`
	}

	var req s1
	json.NewDecoder(r.Body).Decode(&req)
	p.svc.Publish(req.T, req.C)
}
