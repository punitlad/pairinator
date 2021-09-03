package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"pairinator/databases"
	"pairinator/models"
)

func rotateCmd(md databases.MemberDatabase) *cobra.Command {
	return &cobra.Command{
		Use:   "rotate",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			stair, err := models.NewPairStair(md.GetAll())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			stair.Rotate()
			for _, member := range stair.Members {
				md.Update(member)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(rotateCmd(databases.GetMemberDatabaseInstance()))
}
