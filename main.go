package main

import (
	"os"

	"github.com/dev-frame/spograph/generate"
)

func main() {
	rootCMD := generate.GenGraphCommand()
	if err := rootCMD.Execute(); err != nil {
		rootCMD.PrintErrln("Error:", err)
		os.Exit(1)
	}
}
