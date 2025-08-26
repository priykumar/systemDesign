package service

import (
	"fmt"

	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/repo"
)

type SubscriberService interface {
	AddSubscriber(string) error
}

type subscriberService struct {
	db repo.Repository
}

func NewSubscriberService(db repo.Repository) SubscriberService {
	return &subscriberService{db}
}

func (s *subscriberService) AddSubscriber(sub string) error {
	if exist := s.db.IsSubPresent(sub); exist {
		return fmt.Errorf("subscriber %s is already present", sub)
	}

	s.db.CreateNewSub(sub)
	return nil
}
