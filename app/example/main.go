package main

import (
	"log"
	
	"example/bootstrap"
)

func main() {
	if bootErr := bootstrap.Boot(); bootErr != nil {
		log.Fatalln(bootErr)
	}
}
