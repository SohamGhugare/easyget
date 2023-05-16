/*
Copyright © 2023 Soham Ghugare soham.ghugare@gmail.com
*/
package cmd

import (
	"easyget/models"
	"easyget/utility"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the registered aliases",
	Long: `This command lists all the registered aliases and their corresponding original urls. This also contains the alias index which can be used for easy installation through 'easyget get' command`,
	Run: func(cmd *cobra.Command, args []string) {

		// Unmarshalling registered aliases json
		var aliases models.Aliases
		utility.ParseJson(&aliases)

		fmt.Println("Registered aliases: ")

		for i:=0; i<len(aliases.Aliases); i++ {
			fmt.Println(i+1, aliases.Aliases[i].Name, ":", aliases.Aliases[i].Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
