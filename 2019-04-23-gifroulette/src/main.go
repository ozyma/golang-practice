package main

import (
	"fmt"
	"log"
)

//gameSession object will hold all of the necessary
type gameSession struct {
	sessionID          string     `json:"sessionID"`
	sessionCookieToken string     `json:"sessionCookieToken"`
	sessionAdminToken  string     `json:"sessionAdminToken"`
	playerCount        int        `json:"players"`
	questions          []question `json:"questions"`
}

type question struct {
	prompt  string `json:"prompt"`
	gifLink string `json:"gifLink"`
}

func main() {
	log.Print("Starting the server")
	log.Print("Checking the gif link storage server")
	log.Print("Gif link server has been validated")

	mySession := gameSession{
		sessionID:          "IDamskljd29ijc",
		sessionCookieToken: "cookiekadsmkl;m",
		sessionAdminToken:  "ADMINu870nu89unfq2980un809u",
		playerCount:        3,
		questions: []question{
			{"testKey", "testVal"},
		},
	}
	fmt.Println(mySession)
	fmt.Println(mySession.questions.question)
}
