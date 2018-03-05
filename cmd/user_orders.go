package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var allOrdersCmd = &cobra.Command{
	Use:   "open",
	Short: "List open current orders for user.",
	Long:  `List open current orders for user.`,
}

func init() {
	allOrdersCmd.RunE = cmdAllOrders
	userCmd.AddCommand(allOrdersCmd)
}

var getAllOrders = func() ([]cobinhood.Order, error) {
	return cobinhoodClient().GetAllOrders()
}

func doAllOrderssCommand(args []string) (string, error) {

	allorders, err := getAllOrders()

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
	for _, order := range allorders {
		t.AddRow(order.Timestamp.String(),
			fmt.Sprintf("%v", order.Id),
			fmt.Sprintf("%v", order.TradingPair),
			fmt.Sprintf("%v", order.State),
			fmt.Sprintf("%v", order.Side),
			fmt.Sprintf("%v", order.Type),
			fmt.Sprintf("%v", order.Price),
			fmt.Sprintf("%v", order.Size),
			fmt.Sprintf("%v", order.Filled),
		)
	}

	return t.ToString(), nil
}

func cmdAllOrders(cmd *cobra.Command, args []string) error {
	response, err := doAllOrderssCommand(args)

	fmt.Println(response)

	return err
}
