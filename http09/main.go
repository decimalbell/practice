package main

import (
	"flag"
	"fmt"
)

var addr = flag.String("addr", ":80", "server addr")
var root = flag.String("root", "./", "server root")

func main() {
	flag.Parse()
	s := NewServer(*addr, *root)
	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
	return
}
