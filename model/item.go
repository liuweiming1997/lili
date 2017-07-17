package model

import "time"

type Item struct {
	// required
	Identifier string `bson:"identifier"`
	Desc       string `bson:"desc"`
	Name       string `bson:"name"`

	Ref string `bson:"ref"`

	Key        string   `bson:"key"`
	KeyHistory []string `bson:"key_history"`

	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
}

func (i *Item) IsValid() bool {
	if i.Identifier == "" || i.Name == "" || i.Desc == "" {
		return false
	}

	return true
}