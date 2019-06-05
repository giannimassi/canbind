package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"os"
)

func main() {
	var udp = pflag.Bool("udp", false, "If true udp is checked instead of tcp")
	var addr = pflag.StringP("address", "a", "", "Address to bind to (may include port)")
	pflag.Parse()
	if *addr == "" {
		pflag.Usage()
		os.Exit(1)
	}

	p := "tcp"
	if *udp {
		p = "udp"
	}

	ln, err := net.Listen(p, *addr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on %s: %v\n", *addr, err)
		os.Exit(1)
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on %s: %s\n", *addr, err)
		os.Exit(1)
	}

	fmt.Printf("%s can be binded\n", *addr)
	os.Exit(0)
}
