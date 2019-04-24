package main

import (
	"fmt"
	"log" 
	//"flag"
	//"os"
	"net/http"
	"github.com"
)

//handler function handles the index visitation of the URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my program!")
	fmt.Println("This is the URL path: ", r.URL.Path[:])
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	fmt.Println("This is the URL path: ", r.URL)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
