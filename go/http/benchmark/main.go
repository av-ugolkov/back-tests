package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	names := make([]string, 0, len(frameworks()))
	for _, f := range frameworks() {
		names = append(names, f.name)
	}

	name := flag.String("framework", "stdlib", "one of: "+strings.Join(names, "|"))
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	var target framework
	for _, f := range frameworks() {
		if f.name == *name {
			target = f
			break
		}
	}
	if target.serve == nil {
		log.Fatalf("unknown framework %q, want one of: %s", *name, strings.Join(names, ", "))
	}

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	stop := target.serve(ln)
	fmt.Printf("%s listening on %s\n", target.name, ln.Addr())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	stop()
}
