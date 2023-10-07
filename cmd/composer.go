/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// composerCmd represents the composer command
var composerCmd = &cobra.Command{
	Use:   "composer",
	Short: "Run Composer scripts as defined in a project's composer.json file.",
	Long: `When a project is configured to use Composer in the .dev-cli.yml file,
this allows for running scripts as defined in a composer.json file. For example:

 "dev-cli composer install"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("composer called")
	},
}

func init() {
	rootCmd.AddCommand(composerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// composerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// composerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
