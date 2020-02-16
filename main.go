package main

import (
	"os"

	"github.com/bskim45/dtags/cmd"
)

func main() {
	if err := cmd.Execute(os.Args[1:]); err != nil {
		//debug("%+v", err)

		os.Exit(1)
	}
}
