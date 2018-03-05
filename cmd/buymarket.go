package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var buymarketCmd = &cobra.Command{
	Use:   "buymarket",
	Short: "Buy price market.",
	Long:  `Buy price market.`,
}

func init() {
	buymarketCmd.RunE = cmdBuyMarket
	mainCmd.AddCommand(buymarketCmd)
}

var getBuyMarket = func(order cobinhood.PlaceOrder) (cobinhood.Order, error) {
	return cobinhoodClient().PlaceOrder(order)
}

func doBuyMarketCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("expected 'buymarket  TRADING_PAIR  SIZE'\n" +
			"TRADING_PAIR SIZE are required arguments for buymarket command\n" +
			"See 'cobinhood buymarket -h' for help and example."))
	}

	TradingPairId := args[0]
	size := args[1]

	placeorder := cobinhood.PlaceOrder{
		TradingPairId: TradingPairId,
		Side:            "bid",
		Type:            "market",
		Size:            size,
	}
	order, err := getBuyMarket(placeorder)

	if err != nil {
		return "", stdcli.ExitError(err)
	}
	t := stdcli.NewTable("DATE",
		"ID",
		"TRADING_PAIR",
		"STATE",
		"SIDE",
		"TYPE",
		"PRICE",
		"SIZE",
		"FILLED")

	t.AddRow(order.Timestamp.StringHour(),
		fmt.Sprintf("%v", order.Id),
		fmt.Sprintf("%v", order.TradingPair),
		fmt.Sprintf("%v", order.State),
		fmt.Sprintf("%v", order.Side),
		fmt.Sprintf("%v", order.Type),
		fmt.Sprintf("%v", order.Price),
		fmt.Sprintf("%v", order.Size),
		fmt.Sprintf("%v", order.Filled),
	)

	return t.ToString(), nil
}

func cmdBuyMarket(cmd *cobra.Command, args []string) error {
	response, err := doBuyMarketCommand(args)

	fmt.Println(response)

	return err
}
