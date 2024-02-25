package models

type Player struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

var Players []Player
