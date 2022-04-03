package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pairinator/databases"
	"pairinator/models"
)

func addCmd(database databases.MemberDatabase) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new person that can be paired with",
		Run: func(cmd *cobra.Command, args []string) {
			database.Add(models.NewMember(args[0]))
			fmt.Printf("Successfully added pair member %v.\n", args[0])
		},
	}
}

func init() {
	rootCmd.AddCommand(addCmd(databases.GetMemberDatabaseInstance("pairinator.db")))
}
