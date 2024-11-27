package main

import (
	"fmt"
	"github.com/yoricya/bconf-go/bconf"
	"io"
	"log"
	"os"
)

func main() {
	f, e := os.Open("conf.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	b, e := io.ReadAll(f)
	if e != nil {
		log.Fatal(e)
	}

	confStr := string(b)

	config, e := bconf.Parse(confStr)
	if e != nil {
		log.Fatal(e)
	}

	for i := range config {
		fmt.Println("Field:", config[i])
	}
}
