package commands

import (
	"github.com/spf13/cobra"
	"fmt"
)

var userCmd = &cobra.Command{
	Use:   "orders",
	Short: "Get orders from user",
	Long:  `Get orders from user`,
}

func init() {
	userCmd.RunE = cmdUser
	mainCmd.AddCommand(userCmd)
}

func cmdUser(cmd *cobra.Command, args []string) error {

	fmt.Println( "You must specify the type of resource to get. Valid resource types include:\n\n"+
		"   * history\n"+
		"   * info\n"+
		"   * open\n"+
		"   * orderTrades\n"+
		"See 'cobinhood   orders -h' for help and examples.")

	return nil
}
