package app

import (
  "github.com/spf13/cobra"
  navy "bitbucket.org/navy-project/navyapi/client"
  "fmt"
  "io/ioutil"
)

func launchCommand() *cobra.Command {
  return &cobra.Command{
    Use: "launch <name> <manifest.yml>",
    Short: "launch convoy on the cluster",
    Run: launch,
  }
}

func launch(c *cobra.Command, args []string) {
  if len(args) != 2 {
    c.Usage()
    return
  }
  name := args[0]
  manifest_file := args[1]
  manifest, err := readContents(manifest_file)
  if err != nil {
    fmt.Println(err)
    return
  }
  client := navy.NewClient("http://localhost:4040")
  err = client.CreateConvoy(name, manifest)
  if err != nil {
    fmt.Println(err)
    return
  }
}


func readContents(file string) (string, error) {
  bytes, err := ioutil.ReadFile(file)
  if err != nil {
    return "", err
  }
  return string(bytes), nil
}
