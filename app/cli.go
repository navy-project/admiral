package app

import (
  "github.com/spf13/cobra"
)

func BuildApp() *cobra.Command {
  a := &cobra.Command{}
  a.Use = "admiral"
  a.Short = "Command line tool for interacting with the Navy cluster"
  a.AddCommand(launchCommand())
  a.AddCommand(destroyCommand())
  //a.AddCommand(remoteCommand())
  return a
}


func remoteCommand() *cobra.Command {
  return &cobra.Command{
    Use: "remote",
    Short: "specify a named remote cluster",
  }
}
