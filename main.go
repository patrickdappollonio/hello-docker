package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

var (
	port    = ":8000"
	defname = ""
)

func init() {
	// If a port is specified, then use that port
	if v := strings.TrimSpace(os.Getenv("PORT")); v != "" {
		port = fmt.Sprintf(":%v", v)
	}

	// If an environment name is specified, use that instead
	if v := strings.TrimSpace(os.Getenv("NAME")); v != "" {
		defname = v
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	var name string

	// If there's a default name, use it, else use the hostname
	if defname != "" {
		name = defname
	} else {
		name, _ = os.Hostname()
	}

	// If nothing was found, still print "Docker"
	if name == "" {
		name = "Docker"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Handle accessing "/"
	http.HandleFunc("/", hello)

	// Graceful shutdown
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill)

	// Wait signal
	close := make(chan bool, 1)

	// Create a server
	srv := &http.Server{Addr: port}

	// Execute the server
	go func() {
		log.Printf("Starting HTTP Server. Listening at %q", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Println(err.Error())
			} else {
				log.Println("Server closed. Bye!")
			}
			close <- true
		}
	}()

	// Check for a closing signal
	go func() {
		<-sigquit
		log.Printf("Gracefully shutting down server...")

		if err := srv.Shutdown(nil); err != nil {
			log.Println("Unable to shut down server: " + err.Error())
			close <- true
		}
	}()

	<
	
