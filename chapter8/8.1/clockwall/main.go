package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

/*
After starting clock server, this program can be executed by this sample command:
./clockwall US/Eastern=localhost:8010 Asia/Tokyo=localhost:8020
*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("program must have at least 2 arguments")
	}

	d := newClockwallDisplay()

	for _, arg := range os.Args[1:] {
		splits := strings.Split(arg, "=")
		if len(splits) != 2 {
			log.Fatal("arg has wrong format")
		}

		tz := splits[0]
		host := splits[1]

		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Fatal("invalid timezone: ", err)
		}

		w := &locTimeWriter{
			d:   d,
			loc: loc,
		}

		go writeTimeToClock(host, w)
	}

	select {}
}

func writeTimeToClock(host string, w io.Writer) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if _, err := io.Copy(w, conn); err != nil {
		log.Fatal(err)
	}
}

type locTimeWriter struct {
	d   *clockwallDisplay
	loc *time.Location
}

func (w *locTimeWriter) Write(p []byte) (n int, err error) {
	w.d.update(w.loc, fmt.Sprintf("%s", p))
	w.d.refreshDisplay()

	return len(p), nil
}

type clockwallDisplay struct {
	sync.RWMutex
	timeByLoc map[*time.Location]string
}

func newClockwallDisplay() *clockwallDisplay {
	d := new(clockwallDisplay)
	d.timeByLoc = make(map[*time.Location]string)

	return d
}

func (wr *clockwallDisplay) update(loc *time.Location, t string) {
	wr.Lock()
	wr.timeByLoc[loc] = t
	wr.Unlock()
}

func (wr *clockwallDisplay) refreshDisplay() {
	wr.RLock()
	defer wr.RUnlock()

	// TODO: display time by location in a nicer format
	for l, t := range wr.timeByLoc {
		fmt.Fprintf(os.Stdout, "Time in %s is %s", l.String(), t)
	}
}
