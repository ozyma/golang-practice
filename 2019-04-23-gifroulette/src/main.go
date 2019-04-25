package main

import (
	//"encoding/json"
	"fmt"
	//"log"
	//"net/http"
	"math/rand"
	//"time"
	"strconv"
)

//gameSession object will hold all of the necessary
func generateNewGameSession() {
	x := GameSession{
		SessionID:          strconv.Itoa(rand.Intn(1000)),
		SessionCookieToken: strconv.Itoa(rand.Intn(1000)),
		SessionAdminToken:  strconv.Itoa(rand.Intn(1000)),
		PlayerCount:        rand.Intn(8),
		Questions: []Question{
			{"testKey", "testVal"},
			{"testKey2", "testVal2"},
		},
	}
	fmt.Println("")
	fmt.Println("=========================")
	fmt.Println(x)
	fmt.Println("=========================")
	fmt.Println("")
	

}

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

		for i:=0; i<=1000;i++{
			generateNewGameSession()

		}

}
