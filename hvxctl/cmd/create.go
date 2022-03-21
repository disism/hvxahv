/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create database or services.",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "db":
			if err := cockroach.NewDBAddr().Create(args[1]); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s created!\n", args[1])
		case "svc":
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			command := exec.Command("sh", fmt.Sprintf("%s/../hack/dev/create.sh %s", wd, args[1]))
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			if err := command.Run(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(command)
		default:
			fmt.Println("Execution failed, please check the command.")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}