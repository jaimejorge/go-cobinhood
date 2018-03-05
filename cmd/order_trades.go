package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var orderTradesCmd = &cobra.Command{
	Use:   "orderTrades",
	Short: "Get all trades originating from the specific order.",
	Long:  `Get all trades originating from the specific order.`,
}

func init() {
	orderTradesCmd.RunE = cmdOrderTrades
	userCmd.AddCommand(orderTradesCmd)
}

var getOrderTrades = func(trading_id string) (cobinhood.Trades, error) {
	return cobinhoodClient().GetOrderTrades(trading_id)
}

func doOrderTradesCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("Trading-id is required. ex:'COB-BTC'"))
	}

	orderTrades, err := getOrderTrades(args[0])

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("DATE",
		"PRICE",
		"SIZE")
	for _, trade := range orderTrades {
		color := "32"
		if trade.MakerSide == "bid" {
			color = "31"
		}
		t.AddRow(trade.Timestamp.String(),
			fmt.Sprintf("\033[%vm%v\033[0m", color, trade.Price),
			fmt.Sprintf("%v", trade.Size),
		)
	}

	return t.ToString(), nil
}

func cmdOrderTrades(cmd *cobra.Command, args []string) error {
	response, err := doOrderTradesCommand(args)

	fmt.Println(response)

	return err
}
