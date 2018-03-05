package main

import (
	"github.com/jaimejorge/go-cobinhood/cmd"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	commands.Execute(os.Stdin, os.Stdout, os.Stderr)
}
