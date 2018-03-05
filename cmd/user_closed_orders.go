package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var historyOrdersCmd = &cobra.Command{
	Use:   "history",
	Short: "Returns order history for the current user.",
	Long:  `Returns order history for the current user.`,
}

func init() {
	historyOrdersCmd.RunE = cmdHistoryOrders
	userCmd.AddCommand(historyOrdersCmd)
}

var getHistoryOrders = func(TradingPair []string, limit string) ([]cobinhood.Order, error) {
	return cobinhoodClient().GetHistoryOrders(TradingPair, limit)
}

func doHistoryOrderssCommand(args []string) (string, error) {

	historyorders, err := getHistoryOrders(args, "50")

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
	for _, order := range historyorders {
		color := "32"
		if order.Side == "bid" {
			color = "31"
		}
		if order.State == "cancelled" {
			color = "39"
		}
		t.AddRow(order.Timestamp.String(),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Id),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.TradingPair),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.State),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Side),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Type),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Price),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Size),
			fmt.Sprintf("\033[%vm%v\033[0m", color, order.Filled),
		)
	}

	return t.ToString(), nil
}

func cmdHistoryOrders(cmd *cobra.Command, args []string) error {
	response, err := doHistoryOrderssCommand(args)

	fmt.Println(response)

	return err
}
