package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	env := []string{}
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, "PATH") {
			pair := strings.Split(v, "=")
			env = append(env, fmt.Sprintf("PATH=%s:%s", "~user/foo/bar", pair[1]))
		} else {
			env = append(env, v)
		}
	}
	for _, v := range env {
		println(v)
	}
}
