package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}

func Run() {
	cmd := exec.Command("www")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	go Run()
	print("Run...")
	waitForExit()
}
