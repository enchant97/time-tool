package main

import (
	"fmt"

	"github.com/enchant97/time-tool/cli"
)

// set this during build
var Version = "unknown"

func main() {
	if err := cli.Entrypoint(Version); err != nil {
		fmt.Println(err)
	}
}
