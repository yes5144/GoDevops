package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yes5144/GoDevops/GoDevops_api/routers"
)

func main() {
	log.Println("hello")
	r := routers.InitRouter()
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", 8080),
		Handler:           r,
		ReadHeaderTimeout: 60,
		WriteTimeout:      60,
		MaxHeaderBytes:    1 << 20,
	}
	log.Println("http_port: ", 8080)
	s.ListenAndServe()
}
