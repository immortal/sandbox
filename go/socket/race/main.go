package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	count uint
	flag  bool
	quit  chan struct{}
	time  time.Time
	wg    sync.WaitGroup
}

func (c *Counter) Start() {
	c.count = 1
	c.time = time.Now()
	c.flag = true
}

func (c *Counter) Listen() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/", c.HandleStatus)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		log.Println(srv.ListenAndServe())
	}()
	go func(quit chan struct{}) {
		<-quit
		if err := srv.Close(); err != nil {
			log.Printf("HTTP error: %v", err)
		}
	}(c.quit)
}

func (c *Counter) HandleStatus(w http.ResponseWriter, r *http.Request) {
	c.RLock()
	defer c.RUnlock()
	status := struct {
		Count uint   `json:"count"`
		Flag  bool   `json:"flag"`
		Time  string `json:"time"`
	}{
		Count: c.count,
		Time:  c.time.UTC().Format(time.RFC3339),
		Flag:  c.flag,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Println(err)
	}

}

func main() {
	c := &Counter{
		quit: make(chan struct{}),
	}
	c.Start()
	c.Listen()
	timeout := time.After(time.Minute)
	for {
		select {
		case <-time.After(time.Millisecond):
			c.Lock()
			c.count += 1
			c.flag = !c.flag
			c.Unlock()
		case <-timeout:
			close(c.quit)
			c.wg.Wait()
			return
		}
	}
}
