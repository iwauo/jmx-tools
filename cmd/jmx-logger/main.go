package main

import (
	"flag"
	"log"
	"os"

	"github.com/iwauo/jmx-tools/jmxclient"
)

func main() {
	filepath := flag.String("f", "", "JMX watch target definition file")
	flag.Parse()
	if *filepath == "" {
		flag.Usage()
		os.Exit(1)
	}
	config, err := jmxclient.GetConfig(*filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = jmxclient.Start(*config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
