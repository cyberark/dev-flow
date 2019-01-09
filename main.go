package main

import (
	"log"
	
	"github.com/cyberark/dev-flow/cmd"
)

func main() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Llongfile)
	
	cmd.Execute()
}
