package main

import (
	"flag"
	"io"
	"log"
)

func init() {
	var quiet bool
	flag.BoolVar(&quiet, "q", false, "disable log")
	flag.Parse()
	if quiet {
		log.SetOutput(io.Discard)
	}
}

func main() {
	write(process(load(("./..."))))
}
