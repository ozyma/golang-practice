package main

import (
	"encoding/json"
	//"fmt"
	//"log"
	"net/http"
)

//gameSession object will hold all of the necessary
type GameSession struct {
	SessionID          string     `json:"sessionID"`
	SessionCookieToken string     `json:"sessionCookieToken"`
	SessionAdminToken  string     `json:"sessionAdminToken"`
	PlayerCount        int32      `json:"playerCount,string"`
	Questions          []Question `json:"questions"`
}

type Question struct {
	Prompt  string `json:"prompt"`
	GifLink string `json:"gifLink"`
}


//function is used for testing purposes to marshal JSON responses back and
//forth
func indexhandler(w http.ResponseWriter, r *http.Request) {
	mySession := GameSession{
		SessionID:          "IDamskljd29ijc",
		SessionCookieToken: "cookiekadsmkl;m",
		SessionAdminToken:  "ADMINu870nu89unfq2980un809u",
		PlayerCount:        3,
		Questions: []Question{
			{"testKey", "testVal"},
			{"testKey2", "testVal2"},
		},
	}
	var buf GameSession
	json.NewEncoder(w).Encode(mySession)
	
}

func main() {
	http.HandleFunc("/", indexhandler)
	http.ListenAndServe(":8080", nil)
}
