package main

import (
	"fmt"
	"io"
)

func main() {
	var (
		r *io.PipeReader
		w *io.PipeWriter
	)
	fmt.Printf("r = %+v\n", r)
	fmt.Printf("w = %+v\n", w)
	if w != nil {
		w.Close()
	}

}
