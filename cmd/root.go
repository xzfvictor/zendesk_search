/*
Copyright © 2020 Victor Xian

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
	"os"
	"encoding/json"
	gojsonq "github.com/thedevsaddam/gojsonq"
	"github.com/elgs/jsonql"
)

type Users []struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

type Organizations []struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}


type Tickets []struct {
	ID             string	   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zendesk_search -f FILENAME -k KEYNAME -d DATA",
	Short: "Search a key-value pair in a given json file",
	Long: `zendesk_search is a simple CLI tool that can query a key-value pair on a json file`,

	Run: func(cmd *cobra.Command, args []string) {
		filename, _:= cmd.Flags().GetString("filename")
		key, _:= cmd.Flags().GetString("key")
		data, _:= cmd.Flags().GetString("data")
 		if filename == "" || key == "" || data == "" {
			fmt.Println("Search a json file: zendesk_search -f FILENAME -k KEY -d \"DATA\"\n")
			fmt.Println("In Linux/Mac:\n" + "zendesk_search -f users.json -k _id -d \"2\"\n" + "zendesk_search -f tickets.json -k tags -d \"Ohio\"\n" + "zendesk_search -f organizations.json -k name -d \"Xylar\"\n")
			fmt.Println("In Windows:\n" + "zendesk_search.exe -f users.json -k _id -d \"2\"\n" + "zendesk_search.exe -f tickets.json -k tags -d \"Ohio\"\n" + "zendesk_search.exe -f organizations.json -k name -d \"Xylar\"\n")
			fmt.Println("\nList searchable keys:\n" + "zendesk_search checkfile -f FILENAME\n" + "or\n" + "zendesk_search.exe checkfile -f FILENAME\n")
			fmt.Println("Get more help:\n" + "zendesk_search -h\n" + "or\n" + "zendesk_search.exe -h")
		} else {
			switch filename {
			case "users.json":
				searchUsers(filename,key,data)
			case "tickets.json":
				searchTickets(filename,key,data)
			case "organizations.json":
				searchOrganization(filename,key,data)
			default:
				fmt.Println("wrong filename?")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
	rootCmd.Flags().StringP("filename", "f", "", "A json file name: users.json, tickets.json, organizations.json")
	rootCmd.Flags().StringP("key", "k", "", "A key in the json data. Use \"zendesk_search checkfile\" to see all searchable keys.")
	rootCmd.Flags().StringP("data", "d", "", "Data value to be searched")
}

func searchUsers(filename, key, data string) {
	jsonString := gojsonq.New().File("./users.json").Get()

	var searchString string

	if key == "tags" || key == "domain_names" {
		searchString = key + " contains '" + data + "'"
	} else if key == "_id" || key == "organization_id" {
		searchString = key + "=" + data
	} else {
		searchString = key + " = '" + data + "'"
	}

	res, _ := jsonql.NewQuery(jsonString).Query(searchString)

	b, _ := json.Marshal(res)

	var users Users
	json.Unmarshal(b, &users)
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v:			%v\n", "ID", users[i].ID)
		fmt.Printf("%v:			%v\n", "URL", users[i].URL)
		fmt.Printf("%v:		%v\n", "External ID", users[i].ExternalID)
		fmt.Printf("%v:			%v\n", "Name", users[i].Name)
		fmt.Printf("%v:			%v\n", "Alias", users[i].Alias)
		fmt.Printf("%v:		%v\n", "Verified", users[i].Verified)
		fmt.Printf("%v:			%v\n", "Shared", users[i].Shared)
		fmt.Printf("%v:			%v\n", "Locale", users[i].Locale)
		fmt.Printf("%v:		%v\n", "Timezone", users[i].Timezone)
		fmt.Printf("%v:		%v\n", "Last Login At", users[i].LastLoginAt)
		fmt.Printf("%v:			%v\n", "Email", users[i].Email)
		fmt.Printf("%v:			%v\n", "Phone", users[i].Phone)
		fmt.Printf("%v:		%v\n", "Signature", users[i].Signature)
		fmt.Printf("%v:	%v\n", "Organization ID", users[i].OrganizationID)
		fmt.Printf("%v:			%v\n", "Tags", users[i].Tags)
		fmt.Printf("%v:		%v\n", "Suspended", users[i].Suspended)
		fmt.Printf("%v:			%v\n", "Role", users[i].Role)
		fmt.Println("\n")
	}

}

func searchTickets(filename, key, data string) {
	jsonString := gojsonq.New().File("./tickets.json").Get()

	var searchString string

	if key == "tags" || key == "domain_names" {
		searchString = key + " contains '" + data + "'"
	} else if key == "submitter_id" || key == "assignee_id" || key == "organization_id" {
		searchString = key + "=" + data
	} else {
		searchString = key + " = '" + data + "'"
	}
	res, _ := jsonql.NewQuery(jsonString).Query(searchString)

	b, _ := json.Marshal(res)

	var tickets Tickets
	json.Unmarshal(b, &tickets)
	for i := 0; i < len(tickets); i++ {
		fmt.Printf("%v:			%v\n", "ID", tickets[i].ID)
		fmt.Printf("%v:			%v\n", "URL", tickets[i].URL)
		fmt.Printf("%v:		%v\n", "External ID", tickets[i].ExternalID)
		fmt.Printf("%v:		%v\n", "Created At", tickets[i].CreatedAt)
		fmt.Printf("%v:			%v\n", "Type", tickets[i].Type)
		fmt.Printf("%v:		%v\n", "Subject", tickets[i].Subject)
		fmt.Printf("%v:		%v\n", "Description", tickets[i].Description)
		fmt.Printf("%v:		%v\n", "Priority", tickets[i].Priority)
		fmt.Printf("%v:			%v\n", "Status", tickets[i].Status)
		fmt.Printf("%v:		%v\n", "Submitter ID", tickets[i].SubmitterID)
		fmt.Printf("%v:		%v\n", "Assignee ID", tickets[i].AssigneeID)
		fmt.Printf("%v:	%v\n", "Organization ID", tickets[i].OrganizationID)
		fmt.Printf("%v:			%v\n", "Tags", tickets[i].Tags)
		fmt.Printf("%v:		%v\n", "Has Incidents", tickets[i].HasIncidents)
		fmt.Printf("%v:			%v\n", "Due At", tickets[i].DueAt)
		fmt.Printf("%v:			%v\n", "Via", tickets[i].Via)
		fmt.Println("\n")
	}
}

func searchOrganization(filename, key, data string) {
	jsonString := gojsonq.New().File("./organizations.json").Get()

	var searchString string

 	if key == "tags" || key == "domain_names" {
		searchString = key + " contains '" + data + "'"
	} else if key == "_id" {
		searchString = key + "=" + data
	} else {
		searchString = key + " = '" + data + "'"
	} 
	
	res, _ := jsonql.NewQuery(jsonString).Query(searchString)

	b, _ := json.Marshal(res)

	var organizations Organizations
	json.Unmarshal(b, &organizations)
	for i := 0; i < len(organizations); i++ {
		fmt.Printf("%v:			%v\n", "ID", organizations[i].ID)
		fmt.Printf("%v:			%v\n", "URL", organizations[i].URL)
		fmt.Printf("%v:		%v\n", "External ID", organizations[i].ExternalID)
		fmt.Printf("%v:			%v\n", "Name", organizations[i].Name)
		fmt.Printf("%v:		%v\n", "Domain Names", organizations[i].DomainNames)
		fmt.Printf("%v:		%v\n", "Created At", organizations[i].CreatedAt)
		fmt.Printf("%v:		%v\n", "Details", organizations[i].Details)
		fmt.Printf("%v:		%v\n", "Shared Tickets", organizations[i].SharedTickets)
		fmt.Printf("%v:			%v\n", "Tags", organizations[i].Tags)
		fmt.Println("\n")
	}
}