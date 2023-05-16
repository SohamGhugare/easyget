/*
Copyright © 2023 Soham Ghugare soham.ghugare@gmail.com
*/
package cmd

import (
	"easyget/models"
	"easyget/utility"
	"log"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Install a dependency",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1{
			log.Fatal("Provide the dependency name to add to your project.\nUse the --help flag for comprehensive help on how to use this tool")
		}
		dependency := args[0]
		log.Println("Attempting to get", dependency)

		var aliases models.Aliases
		utility.ParseJson(&aliases)

		for _, alias := range aliases.Aliases {
			if alias.Name == dependency{
				utility.ExecShell(alias.Url)
				log.Printf("Successfully added dependency %v", dependency)
				return
			}
		} 
		log.Fatalf("Dependency %v not found.", dependency)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
