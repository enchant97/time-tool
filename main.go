package main

import (
	_ "embed"
	"fmt"
	"github.com/enchant97/time-tool/cli"
	_ "time/tzdata"
)

//go:embed doc.md
var docFileContent string

// set this during build
var Version = "unknown"

func main() {
	if err := cli.Entrypoint(Version, docFileContent); err != nil {
		fmt.Println(err)
	}
}
