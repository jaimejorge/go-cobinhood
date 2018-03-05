package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "Returns ticker for specified trading pair.",
	Long:  `Returns ticker for specified trading pair.`,
}

func init() {
	tickerCmd.RunE = cmdTicker
	marketCmd.AddCommand(tickerCmd)
}

var getTicker = func(trading_id string) (cobinhood.Ticker, error) {
	return cobinhoodClient().GetTicker(trading_id)
}

func doTickerCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("Trading-id is required. ex:'UTNP-ETH'"))
	}

	ticker, err := getTicker(args[0])

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("Last_trade_price",
		"timestamp",
		"24h_high",
		"24h_low",
		"24h_open",
		"24h_volume",
		"last_trade_price",
		"highest_bid",
		"lowest_ask")

	t.AddRow(ticker.LastTradePrice,
		fmt.Sprintf("%v", ticker.Timestamp),
		fmt.Sprintf("%v", ticker.H24hhigh),
		fmt.Sprintf("%v", ticker.H24hlow),
		fmt.Sprintf("%v", ticker.H24hopen),
		fmt.Sprintf("%v", ticker.H24hvolume),
		fmt.Sprintf("%v", ticker.LastTradePrice),
		fmt.Sprintf("%v", ticker.HighestBid),
		fmt.Sprintf("%v", ticker.LowestAsk))

	return t.ToString(), nil
}

func cmdTicker(cmd *cobra.Command, args []string) error {
	response, err := doTickerCommand(args)

	fmt.Println(response)

	return err
}
