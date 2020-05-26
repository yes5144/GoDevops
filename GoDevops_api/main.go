package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yes5144/GoDevops/GoDevops_api/models"
	"github.com/yes5144/GoDevops/GoDevops_api/routers"
	"github.com/yes5144/GoDevops/GoDevops_api/utils"
)

func main() {
	// initConfig
	utils.InitConfig()
	// initDB
	db := models.InitDB()
	defer db.Close()
	// initRouter
	r := routers.InitRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
		// ReadHeaderTimeout: 60,
		// WriteTimeout:      60,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("http_port: ", 8080)
	s.ListenAndServe()
}
