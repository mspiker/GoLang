package main 

import (
	"time"
	"log"
)

type Profile struct {
	Name		string		`json: "username"`
	Password	string		`json: "password"`
	Age			int			'json: "age"`
	LastUpdated	time.Time
}

func GetDB() string {
	return "MongoDB"
}

func GetProfiles() []Profile {
	session, err := mgo.Dial("")
}

