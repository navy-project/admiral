package main

import (
  "bitbucket.org/navy-project/admiral/app"
)

func main() {
  cli := app.BuildApp()
  cli.Execute()
}
