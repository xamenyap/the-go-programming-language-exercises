package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	text := make(chan string, 10)
	go func() {
		for {
			select {
			case tx := <-text:
				go echo(c, tx, 1*time.Second)
			case <-time.After(10 * time.Second):
				fmt.Fprintln(c, "Closing connection because of time out")
				c.Close()
				return
			}
		}
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		text <- input.Text()
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
