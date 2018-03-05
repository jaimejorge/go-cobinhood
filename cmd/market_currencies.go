package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var currenciesCmd = &cobra.Command{
	Use:   "currencies",
	Short: "Returns info for all currencies available for trade.",
	Long:  `Returns info for all currencies available for trade.`,
}

func init() {
	currenciesCmd.RunE = cmdCurrencies
	marketCmd.AddCommand(currenciesCmd)
}

var getCurrencies = func() (cobinhood.Currencies, error) {
	return cobinhoodClient().Getcurrencies()
}

func doCurrenciesCommand(args []string) (string, error) {
	if len(args) > 0 {
		return "", stdcli.ExitError(fmt.Errorf("`info` does not take arguments."))
	}

	currencies, err := getCurrencies()

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("CURRENCY", "NAME", "MIN_UNIT", "DEPOSIT_FEE", "WITHDRAWAL_FEE")

	for _, currency := range currencies {

		t.AddRow(currency.Currency,
			fmt.Sprintf("%v",
				currency.Name),
					fmt.Sprintf("%v", currency.MinUnit),
						fmt.Sprintf("%v", currency.DepositFee),
							fmt.Sprintf("%v", currency.WithdrawalFee))
	}

	return t.ToString(), nil
}

func cmdCurrencies(cmd *cobra.Command, args []string) error {
	response, err := doCurrenciesCommand(args)

	fmt.Println(response)

	return err
}
