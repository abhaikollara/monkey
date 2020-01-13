package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Println("MonkeyLang v.0.0.1")
	repl.Start(os.Stdin, os.Stdout)
}
