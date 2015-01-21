package main

import (
	"github.com/navy-project/admiral/app"
)

func main() {
	cli := app.BuildApp()
	cli.Execute()
}
