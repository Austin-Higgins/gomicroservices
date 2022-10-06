package main

import (
	"github.com/austin-higgins/gomicroservices/handlers"
	"net/http"
	"log"
	"fmt"
	"io"
	"os"
)

func main() { 
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.HandleFunc()

	http.ListenAndServe(":4040", nil)
}