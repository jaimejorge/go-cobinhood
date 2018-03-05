package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var userOrderCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information for a single order.",
	Long:  `Get information for a single order.`,
}

func init() {
	userOrderCmd.RunE = cmdUserOrder
	userCmd.AddCommand(userOrderCmd)
}

var getOrder = func(TradingPair string) (cobinhood.Order, error) {
	return cobinhoodClient().GetOrder(TradingPair)
}

func doOrdersCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("order-id is required. ex:'f58bc99d-e22b-4fca-911a-49b13e8bdea2'"))
	}

	order, err := getOrder(args[0])

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

func cmdUserOrder(cmd *cobra.Command, args []string) error {
	response, err := doOrdersCommand(args)

	fmt.Println(response)

	return err
}
