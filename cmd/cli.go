package commands

import (
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/spf13/cobra"
	"io"
	"os"
	"regexp"
)

const version = "0.1.0"


func Execute(in io.Reader, out, err io.Writer) {
	if _, err := mainCmd.ExecuteC(); err != nil {
		os.Exit(-1)
	}
}

// The main command describes the service and defaults to printing the
// help message.
var mainCmd = &cobra.Command{
	Use:   "cobinhood",
	Short: "Cobinhood is a command line interface to buy, sell and get information in cobinhood.com",
	Long:  `Awesome`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}


var userErrorRegexp = regexp.MustCompile("argument|flag|shorthand")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}


type commandError struct {
	s         string
	userError bool
}

func (u commandError) Error() string {
	return u.s
}

func (u commandError) isUserError() bool {
	return u.userError
}

func cobinhoodClient() *cobinhood.CobinhoodClient {
	host := "api.cobinhood.com"
	apikey := os.Getenv("apikey")
	cl := cobinhood.New(host, apikey)
	return cl
}
