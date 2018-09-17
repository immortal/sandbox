package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/nbari/violetear"
)

type Counter struct {
	count uint32
	flag  bool
	quit  chan struct{}
	wg    sync.WaitGroup
}

type Status struct {
	Count uint32 `json:"count"`
	Flag  bool   `json:"flag"`
}

func (c *Counter) Start() {
	atomic.AddUint32(&c.count, 1)
}

func (c *Counter) Listen() {
	router := violetear.New()
	router.Verbose = false
	router.HandleFunc("/", c.HandleStatus)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		log.Println(srv.ListenAndServe())
	}()
	go func(quit chan struct{}) {
		<-quit
		if err := srv.Close(); err != nil {
			log.Printf("HTTP socket close error: %v", err)
		}
	}(c.quit)
}

func (c *Counter) HandleStatus(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Count: atomic.LoadUint32(&c.count),
		Flag:  c.flag,
	}
	// return status in json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Println(err)
	}

}

func main() {
	counter := &Counter{
		quit: make(chan struct{}),
	}
	counter.Start()
	counter.Listen()
	for {
	}
}
