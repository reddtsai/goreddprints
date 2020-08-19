package main

import (
	"flag"
	"log"
)

func main() {
	name := flag.String("name", "", "who")
	env := flag.String("env", "", "release/debug")
	flag.Parse()
	log.Printf("I am %s.\n", *name)
	log.Printf("%s mode.\n", *env)
}
