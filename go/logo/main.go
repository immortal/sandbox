package main

import (
	"log"
	"strconv"
)

const logo = "2B55"

func main() {
	i, err := strconv.ParseInt(logo, 16, 32)
	if err != nil {
		panic(err)
	}
	log.Printf("logo = %c\n", rune(i))
}
