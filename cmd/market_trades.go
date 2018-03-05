package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var tradesCmd = &cobra.Command{
	Use:   "trades",
	Short: "Returns most recent trades for the specified trading pair.",
	Long:  `Returns most recent trades for the specified trading pair.`,
}

func init() {
	tradesCmd.RunE = cmdTrades
	marketCmd.AddCommand(tradesCmd)
}

var getTrades = func(trading_id string) (cobinhood.Trades, error) {
	return cobinhoodClient().GetTrades(trading_id)
}

func doTradesCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("Trading-id is required. ex:'COB-BTC'"))
	}

	trades, err := getTrades(args[0])

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("DATE",
		"PRICE",
		"SIZE")
	for _, trade := range trades {
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

func cmdTrades(cmd *cobra.Command, args []string) error {
	response, err := doTradesCommand(args)

	fmt.Println(response)

	return err
}
