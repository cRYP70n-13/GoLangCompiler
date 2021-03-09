package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/goCompiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is cRYP70N's programming language!\n", user.Username)
	fmt.Printf("Feel free to type in your code!\n")
	repl.Start(os.Stdin, os.Stdout)
}
