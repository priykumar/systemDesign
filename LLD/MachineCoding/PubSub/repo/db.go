package repo

import (
	"fmt"
)

type Repository interface {
	IsTopicPresent(string) bool
	CreateNewTopic(string) error
	AddSubsToTopic(string, string) error

	IsSubPresent(string) bool
	CreateNewSub(string) error

	IsPubPresent(string) bool
	CreateNewPub(string) error

	GetSubsList(string) []string
}

type repository struct {
	Topic_Subs map[string][]string
	Subs_Id    map[string]bool
	Pubs_Id    map[string]bool
}

func NewRepository() Repository {
	return &repository{
		Topic_Subs: make(map[string][]string),
		Subs_Id:    make(map[string]bool),
		Pubs_Id:    make(map[string]bool),
	}
}

func (r *repository) IsTopicPresent(tName string) bool {
	if _, exist := r.Topic_Subs[tName]; exist {
		return true
	}

	return false
}

func (r *repository) CreateNewTopic(tName string) error {
	r.Topic_Subs[tName] = make([]string, 0)
	fmt.Println("Topic", tName, "is sucessfully added")
	return nil
}

func (r *repository) AddSubsToTopic(tName string, sName string) error {
	r.Topic_Subs[tName] = append(r.Topic_Subs[tName], sName)
	fmt.Println("Subscriber", sName, "has subscriber to topic", tName)
	return nil
}

func (r *repository) IsSubPresent(sName string) bool {
	if _, exist := r.Subs_Id[sName]; exist {
		return true
	}

	return false
}

func (r *repository) CreateNewSub(sName string) error {
	r.Subs_Id[sName] = true
	fmt.Println("Subscriber", sName, "is sucessfully added")
	return nil
}

func (r *repository) IsPubPresent(sName string) bool {
	if _, exist := r.Pubs_Id[sName]; exist {
		return true
	}

	return false
}

func (r *repository) CreateNewPub(sName string) error {
	r.Pubs_Id[sName] = true
	fmt.Println("Publisher", sName, "is sucessfully added")
	return nil
}

func (r *repository) GetSubsList(tName string) []string {
	return r.Topic_Subs[tName]
}
