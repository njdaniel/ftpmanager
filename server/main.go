package main

import (
	"net"
	"log"
	"flag"
	"fmt"
	//"bufio"
)

type conn struct {
	rw 				net.Conn
	dataHostPort	string
	prevCmd			string
	pasvListener	net.Listener
	cmdErr			error
	binary			bool
}

func NewConn(cmdConn net.Conn) *conn {
	return &conn{rw: cmdConn}
}

func (c *conn) run()  {
	fmt.Fprintln("220 Ready.")
	//s := bufio.NewScanner(c.rw)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "listen port")

	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Opening main listener")
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Print("Accepting new connection:", err)
		}
		go NewConn(c).run()
	}
}


