package main

import (
	"os"
	"fmt"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"net/http"
	"time"
	"log"
	"github.com/maxdevelopment/bank-api-demo/routes"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	addr, _ := determineListenAddress()

	r := routes.NewRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(r)

	s := &http.Server{
		Addr:           addr,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}