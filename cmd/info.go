package commands

import (
	"fmt"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/jaimejorge/go-cobinhood/pkg/stdcli"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get system information.",
	Long:  `Get system information.`,
}

func init() {
	infoCmd.RunE = cmdInfo
	mainCmd.AddCommand(infoCmd)
}

var getInfo = func() (cobinhood.Info, error) {
	return cobinhoodClient().Getinfo()
}

func doInfoCommand(args []string) (string, error) {
	if len(args) > 0 {
		return "", stdcli.ExitError(fmt.Errorf("`info` does not take arguments."))
	}

	info, err := getInfo()

	if err != nil {
		return "", stdcli.ExitError(err)
	}

	t := stdcli.NewTable("PHASE", "REVISION")

	t.AddRow(info.Phase, fmt.Sprintf("%v", info.Revision))

	return t.ToString(), nil
}

func cmdInfo(cmd *cobra.Command, args []string) error {
	response, err := doInfoCommand(args)

	fmt.Println(response)

	return err
}
