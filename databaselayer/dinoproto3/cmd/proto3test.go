package main

import (
	"flag"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for server, c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
	case "c":
	}
}
