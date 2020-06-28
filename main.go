package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL")
	// path value is now myapp.log
	// level value is ERROR

	flag.Parse() // path and level are overriden if cmd args have been passed

	f, err := os.Open(*path)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n') // will read line including the delimiter
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(strings.TrimSpace(s))
		}

	}
}
