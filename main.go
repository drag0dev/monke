package main

import (
	"monke/repl"
	"os"
)

func main() {
    repl.Start(os.Stdin, os.Stdout)
}
