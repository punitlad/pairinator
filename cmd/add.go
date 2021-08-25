package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pairinator/databases"
	"pairinator/models"
)

func addCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new person that can be paired with",
		Run: func(cmd *cobra.Command, args []string) {
			add(args[0])
		},
	}
}

func add(name string) {
	instance := databases.GetMemberDatabaseInstance()
	instance.Add(models.NewMember(name))
	fmt.Printf("Successfully added pair member %v.\n", name)
}

func init() {
	rootCmd.AddCommand(addCmd())
}
