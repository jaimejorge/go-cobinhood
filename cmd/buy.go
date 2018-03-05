package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var bidCmd = &cobra.Command{
	Use:   "buy",
	Short: "Place a buy order.",
	Long:  `Place a buy order.`,
}

func init() {
	bidCmd.RunE = cmdBid
	mainCmd.AddCommand(bidCmd)
}

var getBid = func(order cobinhood.PlaceOrder) (cobinhood.Order, error) {
	return cobinhoodClient().PlaceOrder(order)
}

func doBidCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("expected 'buy  TRADING_PAIR PRICE SIZE'\n" +
			"TRADING_PAIR PRICE SIZE are required arguments for buy command\n" +
			"See 'cobinhood buy -h' for help and example."))
	}

	TradingPairId := args[0]
	price := args[1]
	size := args[2]

	placeorder := cobinhood.PlaceOrder{
		TradingPairId: TradingPairId,
		Side:            "bid",
		Type:            "limit",
		Price:           price,
		Size:            size,
	}
	order, err := getBid(placeorder)

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

func cmdBid(cmd *cobra.Command, args []string) error {
	response, err := doBidCommand(args)

	fmt.Println(response)

	return err
}
