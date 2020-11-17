package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/j1018813/devtig/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("🦈 %s! This is the Devtig programming language! "+
		"I'd strongly recommend against using it.\n",
		user.Username)

	fmt.Printf("Try it out... But be warned!\n")
	repl.Start(os.Stdin, os.Stdout)
}
