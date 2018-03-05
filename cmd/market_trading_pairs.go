package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var trading_pairsCmd = &cobra.Command{
	Use:   "trading_pairs",
	Short: "Get info for all trading pairs.",
	Long:  `Get info for all trading pairs.`,
}

func init() {
	trading_pairsCmd.RunE = cmdTrading_pairs
	marketCmd.AddCommand(trading_pairsCmd)
}

var getTrading_pairs = func() (cobinhood.TradingPairs, error) {
	return cobinhoodClient().GetTradingPairs()
}

func doTrading_pairsCommand(args []string) (string, error) {
	if len(args) > 0 {
		return "", stdcli.ExitError(fmt.Errorf("`info` does not take arguments."))
	}

	trading_pairs, err := getTrading_pairs()

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("ID", "base_currency_id", "quote_currency_id", "base_max_size", "base_min_size", "quote_increment")

	for _, trading_pair := range trading_pairs {

		t.AddRow(trading_pair.Id,
			fmt.Sprintf("%v", trading_pair.BaseCurrencyId),
				fmt.Sprintf("%v", trading_pair.QuoteCurrencyId),
					fmt.Sprintf("%v", trading_pair.BaseMaxSize),
						fmt.Sprintf("%v", trading_pair.BaseMinSize),
							fmt.Sprintf("%v", trading_pair.QuoteIncrement))
	}

	return t.ToString(), nil
}

func cmdTrading_pairs(cmd *cobra.Command, args []string) error {
	response, err := doTrading_pairsCommand(args)

	fmt.Println(response)

	return err
}
