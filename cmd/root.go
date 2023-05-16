/*
Copyright © 2023 Soham Ghugare soham.ghugare@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "easyget",
	Short: "A easy-to-use replace for vanilla `go get`",
	Long: ``,
	
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


