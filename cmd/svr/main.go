package main

import (
	"context"
	"fmt"
	"github.com/CalebTracey/ultrasound-api/internal/facade"
	"github.com/CalebTracey/ultrasound-api/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const Port = 8000

func main() {
	var wait time.Duration
	defer deathScream()

	service, svcErr := facade.NewUltrasoundService()
	if svcErr != nil {
		log.Panic(svcErr)
	}

	handler := routes.Handler{
		Service: service,
	}

	router := handler.InitializeRoutes()

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", Port),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 120,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	msg := fmt.Sprintf("Listening on Port: %v", Port)
	fmt.Println("\033[36m", msg, "\033[0m")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	_ = srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)

}

func deathScream() {
	if r := recover(); r != nil {
		log.Println(fmt.Errorf("I panicked and am quitting: %v", r))
		log.Println("I should be alerting someone...")
	}
}
