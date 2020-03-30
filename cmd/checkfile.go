/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
)

// checkfileCmd represents the checkfile command
var checkfileCmd = &cobra.Command{
	Use:   "checkfile",
	Short: "Show searchable keys on the given json file",
	Long: `Usage: zendesk_search checkfile -f FILENAME`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _:= cmd.Flags().GetString("filename")
		if filename == "" {
			fmt.Println("Please run the command with -f FILENAME.")
			fmt.Println("\nFor example:\n" + "In Linux/MAC:\n" + "zendesk_search checkfile -f users.json")
			fmt.Println("\nIn Windows:\n" + "zendesk_search.exe checkfile -f users.json")
		} else {
			switch filename {
			case "users.json":
				userdata(filename)
			case "tickets.json":
				ticketdata(filename)
			case "organizations.json":
				organizationdata(filename)
			default:
				fmt.Println("wrong filename?")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkfileCmd)
	checkfileCmd.Flags().StringP("filename", "f", "", "path to the json file name")
}

func userdata(file string) {
	const keys = `_id url external_id name alias created_at active verified shared locale timezone last_login_at email phone signature organization_id tags suspended role`
	fmt.Println("Searchable keys are: \n" + keys)
}

func ticketdata(file string) {
	const keys = `_id url external_id created_at type subject description priority status submitter_id assignee_id organization_id tags has_incidents due_at via`
	fmt.Println("Searchable keys are: \n" + keys)
}

func organizationdata(file string) {
	const keys = `_id url external_id name domain_names created_at details shared_tickets tags`
	fmt.Println("Searchable keys are: \n" + keys)
}
