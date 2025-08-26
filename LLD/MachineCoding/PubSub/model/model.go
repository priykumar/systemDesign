package model

type Topic struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Subs []Subcribers `json:"subs"`
}

type Publisher struct {
	ID   int
	Name string
}

type Subcribers struct {
	ID   int
	Name string
}
