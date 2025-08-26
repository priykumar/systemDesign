package service

import (
	"fmt"

	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/repo"
)

type pubsub struct {
	Topic string
	Ctx   string
}

var Pubsubchan chan pubsub

type PublisherService interface {
	AddPublisher(string) error
	Publish(string, string)
}

type publisherService struct {
	db repo.Repository
}

func NewPublisherService(db repo.Repository) PublisherService {
	Pubsubchan = make(chan pubsub, 2)
	return &publisherService{db}
}

func (p *publisherService) AddPublisher(pname string) error {
	if exist := p.db.IsPubPresent(pname); exist {
		return fmt.Errorf("publisher %s already exixts", pname)
	}

	p.db.CreateNewPub(pname)
	return nil
}

func (p *publisherService) notifySubs(topic, ctx string) {
	subs := p.db.GetSubsList(topic)
	for _, s := range subs {
		fmt.Println("Notifying subscriber", s, "about topic:", topic, "\tcontext:", ctx)
	}
}

func ListenToProducers(p PublisherService) {
	fmt.Println("Listing to producers")
	for r := range Pubsubchan {
		fmt.Println("channel populated")
		go p.(*publisherService).notifySubs(r.Topic, r.Ctx) //notifySubs(r.Topic, r.Ctx)
	}
}

func (p *publisherService) Publish(tName, ctx string) {
	Pubsubchan <- pubsub{Topic: tName, Ctx: ctx}
}
