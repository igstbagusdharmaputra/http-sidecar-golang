package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello!")
}
func main() {
	p := os.Getenv("PORT")
	if p == "" {
		log.Fatal("Port Env Not Define")
	}
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%v", p), nil)
}
