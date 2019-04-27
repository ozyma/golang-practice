package main

import "fmt"

/*GameSession will hold all the necessary information for a single lobby of
players to connect and play a game

*/
type GameSession struct {
	SessionID          string     `json:"SessionID"`
	SessionCookieToken string     `json:"SessionCookieToken"`
	SessionAdminToken  string     `json:"SessionAdminToken"`
	PlayerList         []Player   `json:"PlayerList"`
	PlayerCount        int        `json:"PlayerCount"`
	Questions          []Question `json:"questions"`
}

type Question struct {
	Prompt  string `json:"prompt"`
	GifLink string `json:"gifLink"`
}

type Player struct {
	IP string `json:"IP"`
}

func (r GameSession) PrintGameSession() {
	fmt.Println(r)
}

func (r GameSession) ChangeIDName(){
	r.SessionID = "RogerWilko"
	r.PrintGameSession()
}
