package main

import (
	"log"
	
	"github.com/cyberark/dev-flow/cmd"
)

func main() {
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Llongfile)
	
	cmd.Execute()
}
