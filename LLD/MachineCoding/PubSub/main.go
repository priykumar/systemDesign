package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/controller"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/repo"
	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/service"
)

func main() {
	db := repo.NewRepository()
	tsvc := service.NewTopicService(db)
	ssvc := service.NewSubscriberService(db)
	psvc := service.NewPublisherService(db)
	topic_handler := controller.NewTypeController(tsvc)
	sub_handler := controller.NewSubController(ssvc)
	pub_handler := controller.NewPubController(psvc)

	go service.ListenToProducers(psvc)

	r := mux.NewRouter()
	r.HandleFunc("/addtopic/{topic}", topic_handler.CreateTopicHandler).Methods("POST")
	r.HandleFunc("/subscribetopic", topic_handler.AddSubsHandler).Methods("POST")

	r.HandleFunc("/addsubscriber/{sub}", sub_handler.NewSubHandler).Methods("POST")

	r.HandleFunc("/addpublisher/{sub}", pub_handler.AddPubHandler).Methods("POST")
	r.HandleFunc("/publish", pub_handler.PublishHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
