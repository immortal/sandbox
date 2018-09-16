package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/nbari/violetear"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm catching all\n"))
}

func main() {
	l, err := net.Listen("unix", "/tmp/immortal.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	router := violetear.New()
	router.Verbose = false
	router.HandleFunc("*", catchAll)

	srv := &http.Server{Handler: router}

	closeSocket := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := srv.Close(); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(closeSocket)
	}()

	log.Fatal(srv.Serve(l))

	<-closeSocket
}
