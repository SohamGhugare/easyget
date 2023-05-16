/*
Copyright © 2023 Soham Ghugare soham.ghugare@gmail.com
*/
package cmd

import (
	"easyget/models"
	"easyget/utility"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "Install a dependency",
	Long: `This command goes through the list of all the registered aliases and gets the associated dependency. Equivalent to 'go get [url]' where url is the actual url to the alias dependency`,
	Example: "  easyget get cobra\n  easyget get 1",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1{
			log.Fatal("Provide the dependency name to add to your project.\nUse the --help flag for comprehensive help on how to use this tool")
		}
		dependency := args[0]

		var aliases models.Aliases
		utility.ParseJson(&aliases)

		if dep, err:=strconv.Atoi(dependency); err == nil {
			if dep > len(aliases.Aliases){
				log.Fatal("Invalid index. Use `easyget list` to check the registered aliases.")
				return
			}
			utility.ExecShell(aliases.Aliases[dep-1].Url)
			
		} else {
			for _, alias := range aliases.Aliases {
				if alias.Name == dependency{
					utility.ExecShell(alias.Url)
					return
				}
			} 
			log.Fatalf("Dependency %v not found.", dependency)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
