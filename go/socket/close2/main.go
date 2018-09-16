package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

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

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Println(srv.Serve(l))
		wg.Done()
	}()

	closeSocket := make(chan struct{})

	go func() {
		<-closeSocket
		if err := srv.Close(); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		log.Println("fin")
	}()

	<-time.After(2 * time.Second)
	log.Println("close socket")
	close(closeSocket)
	wg.Wait()
}
