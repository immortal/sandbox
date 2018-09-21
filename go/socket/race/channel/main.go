package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
This code contains data races because the updateFooSlice() goroutine
is changing the same *FooSlice as is being served up by the
http handler without any sync protection.
Should a sync.mutex variable be added or is there a slicker way
to expose internal memory structures to an http handler?
How to reproduce:
    $ go run -race main.go
    $ httpie http://localhost:8080
*/

type Foo struct {
	content string
}

type FooSlice []*Foo

var request chan chan FooSlice

func updateFooSlice(fooSlice FooSlice) {
	t := time.Tick(10 * time.Second)
	for {
		select {
		case <-t:
			foo := &Foo{content: "new"}
			fooSlice[0] = foo
			fooSlice[1] = nil
		case ch := <-request:
			fmt.Println("request...")
			fooSliceCopy := make(FooSlice, len(fooSlice))
			copy(fooSliceCopy, fooSlice)
			ch <- fooSliceCopy
		}
	}
}

func main() {
	request = make(chan chan FooSlice)

	foo1 := &Foo{content: "hey"}
	foo2 := &Foo{content: "yo"}

	fooSlice := FooSlice{foo1, foo2}

	handler := func(w http.ResponseWriter, r *http.Request) {
		response := make(chan FooSlice)
		request <- response
		fooSliceCopy := <-response
		for _, foo := range fooSliceCopy {
			if foo != nil {
				fmt.Fprintf(w, "foo: %v,  time: %v\n", (*foo).content, time.Now().Format(time.RFC3339))
			}
		}
	}
	http.HandleFunc("/", handler)

	go updateFooSlice(fooSlice)
	http.ListenAndServe(":8080", nil)
}
