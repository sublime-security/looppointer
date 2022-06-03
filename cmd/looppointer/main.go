package main

import (
	"github.com/sublime-security/looppointer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(looppointer.Analyzer)
}
