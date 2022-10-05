package main

import (
	"net/http"
	"log"
)


func main(){
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		log.Println("Hello World")
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":4040", nil)
}