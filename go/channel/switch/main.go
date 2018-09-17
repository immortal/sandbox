package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	srv := &http.Server{Addr: ":8080"}

	wg.Add(1)
	go func() {
		log.Println(srv.ListenAndServe())
		wg.Done()
	}()

	quit := make(chan struct{})
	go func() {
		<-quit
		if err := srv.Close(); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		<-quit
		log.Println("just waiting 1")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		<-quit
		log.Println("just waiting 2")
		wg.Done()
	}()

	<-time.After(2 * time.Second)
	close(quit)
	wg.Wait()
}
