package app

import (
	"fmt"

	navy "github.com/navy-project/navyapi/client"
	"github.com/spf13/cobra"
)

func destroyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "destroy <name>",
		Short: "destroy named convoy on the cluster",
		Run:   destroy,
	}
}

func destroy(c *cobra.Command, args []string) {
	if len(args) != 1 {
		c.Usage()
		return
	}
	name := args[0]
	client := navy.NewClient("http://localhost:4040")
	err := client.DeleteConvoy(name)
	if err != nil {
		fmt.Println(err)
		return
	}
}
