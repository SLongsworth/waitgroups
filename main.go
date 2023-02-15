//Filename: main.go
//demonstrating flags
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//set the flags   name of flag, default value of flag, if u dont know how to use the flag
	msg := flag.String("msg", "howdy, stranger!", "the message to display")
	num := flag.Int("num", 1, "how many times to print the message")
	caps := flag.Bool("U", false, "specify whether yo uppercase the stringn (true of false)")

	flag.Parse()
	//check if it should be uppercase before printing it
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	//print the string

	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
