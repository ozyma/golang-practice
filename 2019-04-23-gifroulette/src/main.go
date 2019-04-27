package main

import (
	//"encoding/json"
	//"fmt"
	//"log"
	//"net/http"
	//"math/rand"
	//"time"
	//"strconv"
)

/*func indexhandler(w http.ResponseWriter, r *http.Request) {
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

}*/

func main() {
	/*	http.HandleFunc("/", indexhandler)
		http.ListenAndServe(":8080", nil)*/

	DB := make(map[string]*GameSession)
	
	DB["rustygiraffe223"] = &GameSession{
		SessionID:          "rustygiraffe223",
		SessionCookieToken: "IblPL00sHz7emKLtztel1g==", //Key:tannerandjordan
		SessionAdminToken:  "jwKeRVeWNS2n9N+1pYpFHg==", //Key:tannerandjordanadmin
		PlayerList: []Player{
			{"192.168.1.1"},
		},
		PlayerCount:        3,
		Questions: []Question{
			{"Why is your dog making this face?", "https://media.giphy.com/media/21GCae4djDWtP5soiY/giphy.gif"},
			{"When you notice ________", "https://media.giphy.com/media/lJ0JGfNBrRWJVCRChd/giphy.gif"},
			{"When your text from your Mom says ________","https://media.giphy.com/media/oOxBQwNqGwxeWLDF6A/giphy.gif"},
		},
	}

	//fmt.Println(DB["rustygiraffe223"])
	DB["rustygiraffe223"].ChangeIDName()

}
