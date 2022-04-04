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
			for _, arg := range args {
				database.Add(models.NewMember(arg))
				fmt.Printf("Successfully added pair member %v.\n", arg)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(addCmd(databases.GetMemberDatabaseInstance("pairinator.db")))
}
