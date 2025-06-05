/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// announcmentsCmd represents the announcments command
var announcmentsCmd = &cobra.Command{
	Use:   "announcments ID LOCATION STARTTIME ENDTIME ALLDAY TEAM|TEAM",
	Short: "Interact with announcemnts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("announcments called")
	},
}

func init() {
	rootCmd.AddCommand(announcmentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// announcmentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// announcmentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
