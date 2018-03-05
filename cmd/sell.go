package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "sell",
	Short: "Place a sell order.",
	Long:  `Place a sell order.`,
}

func init() {
	askCmd.RunE = cmdAsk
	mainCmd.AddCommand(askCmd)
}

var getAsk = func(order cobinhood.PlaceOrder) (cobinhood.Order, error) {
	return cobinhoodClient().PlaceOrder(order)
}

func doAskCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("expected 'sell  TRADING_PAIR PRICE SIZE'\n" +
			"TRADING_PAIR PRICE SIZE are required arguments for ask command\n" +
			"See 'cobinhood ask -h' for help and example."))
	}

	TradingPairId := args[0]
	price := args[1]
	size := args[2]

	placeorder := cobinhood.PlaceOrder{
		TradingPairId: TradingPairId,
		Side:            "ask",
		Type:            "limit",
		Price:           price,
		Size:            size,
	}
	order, err := getAsk(placeorder)

	if err != nil {
		return "", stdcli.ExitError(err)
	}
	t := stdcli.NewTable("DATE",
		"ID",
		"trading_pair",
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

func cmdAsk(cmd *cobra.Command, args []string) error {
	response, err := doAskCommand(args)

	fmt.Println(response)

	return err
}
