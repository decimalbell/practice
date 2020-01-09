package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Server struct {
	addr string
	root string
}

func NewServer(addr, root string) *Server {
	if root == "" {
		root = "./"
	}
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	return &Server{
		addr: addr,
		root: root,
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go s.handleConn(conn)
	}

	return nil
}

var errHTML = `<html>
<body>
%s
</body>
</html>
`

func (s *Server) handleConn(conn net.Conn) {
	var err error

	defer func() {
		if err != nil {
			errResp := fmt.Sprintf(errHTML, err)
			conn.Write([]byte(errResp))
		}

		conn.Close()
	}()

	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)
	ss := strings.Split(line, " ")
	if len(ss) != 2 {
		return
	}

	method, path := ss[0], ss[1]
	if !strings.EqualFold(method, "GET") {
		err = fmt.Errorf("http09: method %s not supported", method)
		return
	}

	path = s.root + path
	file, err := os.Open(path)
	if err != nil {
		return
	}
	_, err = io.Copy(conn, file)
	if err != nil {
		return
	}
}
