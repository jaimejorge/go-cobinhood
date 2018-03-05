package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var orderbooksCmd = &cobra.Command{
	Use:   "orderbooks",
	Short: "Get order book for the trading pair containing all asks/bids.",
	Long:  `Get order book for the trading pair containing all asks/bids.`,
}

func init() {
	orderbooksCmd.RunE = cmdOrderbooks
	marketCmd.AddCommand(orderbooksCmd)
}

var getOrderbooks = func(trading_id string) (cobinhood.Orderbook, error) {
	return cobinhoodClient().GetOrderbooks(trading_id)
}

func doOrderbooksCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("Trading-id is required. ex:'UTNP-ETH'"))
	}

	orderbooks, err := getOrderbooks(args[0])

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("PRICE", "COUNT", "SIZE")
	fmt.Sprintf("%v", orderbooks)

	t.AddRow("-------", "-------", "--------")

	for i := len(orderbooks.Asks) - 1; i >= 0; i-- {
		t.AddRow(orderbooks.Asks[i][0], fmt.Sprintf("%v", orderbooks.Asks[i][1]), fmt.Sprintf("%v", orderbooks.Asks[i][2]))
	}

	t.AddRow("---", "---", "---")
	for _, ask := range orderbooks.Bids {
		t.AddRow(ask[0], fmt.Sprintf("%v", ask[1]), fmt.Sprintf("%v", ask[2]))
	}

	return t.ToString(), nil
}

func cmdOrderbooks(cmd *cobra.Command, args []string) error {
	response, err := doOrderbooksCommand(args)

	fmt.Println(response)

	return err
}
