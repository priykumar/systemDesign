package service

import (
	"fmt"

	"github.com/priykumar/systemDesign/LLD/MachineCoding/PubSub/repo"
)

type TopicService interface {
	CreateTopic(topic string) error
	AddSubsToTopic(string, string) error
}

type topicService struct {
	db repo.Repository
}

func NewTopicService(db repo.Repository) TopicService {
	return &topicService{db}
}

func (t *topicService) CreateTopic(topic string) error {
	if exist := t.db.IsTopicPresent(topic); exist {
		return fmt.Errorf("topic name already exist")
	}

	err := t.db.CreateNewTopic(topic)
	if err != nil {
		return err
	}

	return nil
}

func (t *topicService) AddSubsToTopic(topic, subs string) error {
	if exist := t.db.IsTopicPresent(topic); !exist {
		return fmt.Errorf("topic %s doesn't exist", topic)
	}

	if exist := t.db.IsSubPresent(subs); !exist {
		return fmt.Errorf("subscriber %s doesn't exist", subs)
	}

	err := t.db.AddSubsToTopic(topic, subs)
	if err != nil {
		return err
	}

	return nil
}
