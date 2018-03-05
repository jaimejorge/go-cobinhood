package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var CancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel order.",
	Long:  `Cancel order.`,
}

func init() {
	CancelCmd.RunE = cmdCancel
	mainCmd.AddCommand(CancelCmd)
}

var getCancel = func(orderId string) error {
	return cobinhoodClient().CancelOrder(orderId)
}

func doCancelCommand(args []string) (string, error) {

	if len(args) == 0 {
		return "", stdcli.ExitError(fmt.Errorf("expected 'cancel  ORDER_ID'\n" +
			"ORDER_ID are required arguments for cancel command\n" +
			"See 'cobinhood cancel -h' for help and example."))
	}

	orderId := args[0]

	err := getCancel(orderId)

	if err != nil {
		return "", stdcli.ExitError(err)
	}
    msg:=orderId + " order cancelled."
	return msg, nil
}

func cmdCancel(cmd *cobra.Command, args []string) error {
	response, err := doCancelCommand(args)

	fmt.Println(response)

	return err
}
