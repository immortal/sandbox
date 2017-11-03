// +build freebsd netbsd openbsd dragonfly darwin

package testvet

import "sync"

type ScanDir struct {
	dir      string
	services sync.Map
}

func main() {
	s := new(ScanDir)
	s.services.Delete("foo")
}
