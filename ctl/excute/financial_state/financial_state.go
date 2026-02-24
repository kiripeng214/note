package financial_state

import (
	"fmt"

	"github.com/spf13/cobra"
)

func FinancialStateCmd(cmd *cobra.Command, args []string) {
	fmt.Println(FileStr)
}
