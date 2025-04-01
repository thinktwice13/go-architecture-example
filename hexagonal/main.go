package main

import (
	"hex/application"
	"log"
)

func main() {
	log.Fatal(application.BootstrapAndRun())
}
