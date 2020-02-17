package main

import (
	"log"

	"github.com/iwauo/jmx-tools/jmxclient"
)

func main() {
	log.Println("Running...")
	result, err := jmxclient.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
