package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		log.Print("The service is ready to listen and serve.")
		log.Fatal(srv.ListenAndServe())
	}()

	<-stop
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server Exited Properly")
}
