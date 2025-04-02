package main

import (
	"clean/app"
	"clean/bootstrap"
	"log"
)

func main() {
	log.Fatal(bootstrap.RunApplication())
}
