package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pairinator/databases"
)

func viewCmd(md databases.MemberDatabase) *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			all := md.GetAll()

			length := len(all)
			if length == 0 {
				fmt.Println("No members added to pairinator!")
			} else {
				fmt.Println(all)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(viewCmd(databases.GetMemberDatabaseInstance("pairinator.db")))
}
