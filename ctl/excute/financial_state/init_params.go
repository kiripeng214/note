package financial_state

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	rootCmd.Flags().StringVarP(&FileStr, "file", "f", "", "文件地址")
	rootCmd.Flags().StringVarP(&OutStr, "out", "o", "", "输出地址")
}
