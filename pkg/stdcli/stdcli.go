package stdcli

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func ExitError(err error) error {
	return cli.NewExitError(fmt.Sprintf("%s", err.Error()), 1)
}
