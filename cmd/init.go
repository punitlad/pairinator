/*
Copyright © 2021 Punit Lad punitlad@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/buntdb"
)

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initializes pairinator database",
		Long: "Creates a local pairinator database that is used to manage different members of the team.",
		Run: func(cmd *cobra.Command, args []string) {
			buntdb.Open("pairinator.db")
			fmt.Println("Created pairinator database")
		},
	}
}

func init() {
	rootCmd.AddCommand(initCmd())
}
