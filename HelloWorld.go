package main

import (
	"fmt"
	"os"
	"strings"
	)

func main()  {
	who := "world"

	//If there is a command line argument, it will allocate it to who:
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who) //Prints
}
