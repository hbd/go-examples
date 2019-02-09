package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/goodbye", http.HandlerFunc(goodbyeHandler))
	server := &http.Server{
		Addr: "8080",
	}

	// Catch SIGINT to end cleanly and write coverage results to disk.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start the server.
	go func() {
		_ = http.ListenAndServe(":8080", mux)
	}()

	// Ensure a graceful shutdown.
	_ = <-c
	fmt.Println("\nDone")
	server.Shutdown(context.Background())
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!"))
}

func goodbyeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Goodbye :("))
}
