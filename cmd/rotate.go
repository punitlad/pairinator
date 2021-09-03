package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pairinator/databases"
	"pairinator/util"
)

func rotateCmd(md databases.MemberDatabase) *cobra.Command {
	return &cobra.Command{
		Use:   "rotate",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			all := md.GetAll()
			if len(all) < 4 {
				fmt.Println("Not enough members to rotate! Make sure to add more than 3 members.")
			} else {


				var assigned []string

				for _, member := range all {
					if !util.Contains(assigned, member.Name) {
						for _, availablePair := range all {
							if !util.Contains(assigned, availablePair.Name) {
								if !member.HasPairedWith(availablePair.Name) {
									availablePair.AddMemberToPairList(member.Name)
									member.AddMemberToPairList(availablePair.Name)
									assigned = append(assigned, member.Name, availablePair.Name)
									md.Update(member)
									md.Update(availablePair)
									fmt.Printf("Pair Swapped! %s-%s\n", member.Name, availablePair.Name)
									break
								}
							}
						}
					}
				}
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(rotateCmd(databases.GetMemberDatabaseInstance()))
}
