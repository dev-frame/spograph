package main

import (
	"os"

	"github.com/zhufuyi/spograph/generate"
)

func main() {
	rootCMD := generate.GenGraphCommand()
	if err := rootCMD.Execute(); err != nil {
		rootCMD.PrintErrln("Error:", err)
		os.Exit(1)
	}
}
