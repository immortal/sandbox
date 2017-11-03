// +build linux

package testvet

type ScanDir struct {
	dir      string
	services map[string]string
}

func main() {
	s := new(ScanDir)
	delete(s.services, "foo")
}
