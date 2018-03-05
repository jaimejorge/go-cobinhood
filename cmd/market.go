package commands

import (
	"github.com/spf13/cobra"
	"fmt"
)

var marketCmd = &cobra.Command{
	Use:   "market",
	Short: "Get market information,currencies,orderbooks,ticker,trades.",
	Long:  `Get market information,currencies,orderbooks,ticker,trades.`,
}

func init() {
	marketCmd.RunE = cmdMarket
	mainCmd.AddCommand(marketCmd)
}

func cmdMarket(cmd *cobra.Command, args []string) error {
	//response, err := doBidCommand(args)

	fmt.Println( "You must specify the type of resource to get. Valid resource types include:\n\n"+
		"   * currencies\n"+
	"   * orderbooks\n"+
		"   * ticker\n"+
		"   * trades\n"+
		"   * trading_pairs\n"+
			"See 'cobinhood   market -h' for help and examples.")

	return nil
}
