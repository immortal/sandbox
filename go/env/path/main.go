package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("PATH")
	fmt.Println(fmt.Sprintf("PATH=%s:%s", "~user/foo/bar", path))
}
