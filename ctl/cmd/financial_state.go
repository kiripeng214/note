/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ctl/excute/financial_state"

	"github.com/spf13/cobra"
)

// financialStateCmd represents the financialState command
var financialStateCmd = &cobra.Command{
	Use:   "financial_state",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: financial_state.FinancialStateCmd,
}

func init() {
	rootCmd.AddCommand(financialStateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// financialStateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// financialStateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	financial_state.Init(financialStateCmd)

}
