package main

import (
	"net"
	"log"
	"flag"
	"fmt"
	"bufio"
	"strings"
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
	s := bufio.NewScanner(c.rw)
	var cmd string
	var args []string
	for s.Scan() {
		if c.CmdErr() != nil {
			log.Print("err:", fmt.Errorf("command connection: %s", c.CmdErr()))
		}
		fields := strings.Fields(s.Text())
		if len(fields) == 0 {
			continue
		}
		cmd = strings.ToUpper(fields[0])
		args = nil
		if len(fields) > 1 {
			args = fields[1:]
		}
		switch cmd {
		case "LIST":
			c.list(args)
		case "NOOP":
			fmt.Fprintln("200 Ready.")
		case "PASV":
			c.pasv(args)
		case "PORT":
			c.port(args)
		case "QUIT":
			fmt.Fprintln("221 Goodbye.")
			return
		case "RETR":
			c.retr(args)
		case "STOR":
			c.stor(args)
		case "STRU":
			c.stru(args)
		case "SYST":
			fmt.Fprintln("215 UNIX Type: L8")
		case "TYPE":
			c.type_(args)
		case "USER":
			fmt.Fprintln("230 Login successful.")
		default:
			fmt.Fprintln("502 Command not implemented")
		}
	}
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


