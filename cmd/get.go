/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		fmt.Println("Attempting to get", dependency)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
