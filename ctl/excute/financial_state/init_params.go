package financial_state

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	rootCmd.Flags().StringVarP(&FileStr, "file", "f", "", "file path")
	rootCmd.Flags().StringVarP(&OutStr, "out", "o", "", "out path")
}
