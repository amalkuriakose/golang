package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/profile", handleClientProfile)
	log.Println("Server Running on Port 8088...")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
