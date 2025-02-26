package cmd

import (
	"fmt"
	"os"

	"github.com/criblio/scope/internal"
	"github.com/criblio/scope/start"
	"github.com/criblio/scope/util"
	"github.com/spf13/cobra"
)

/* Args Matrix (X disallows)
 *                 force
 * force           -
 */

func getStopUsage() string {
	return `The following actions will be performed on the host and in all relevant containers:
	- Removal of filter files /usr/lib/appscope/scope_filter and /tmp/appscope/scope_filter
	- Detach from all existing scoped processes
	- Removal of etc/profile.d/scope.sh script 
	- Update the relevant service configurations to not LD_PRELOAD libscope if already doing so`
}

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop scoping all scoped processes and services",
	Long: `Stop scoping all processes and services on the host and in all relevant containers.

` + getStopUsage(),
	Example: `  scope stop`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitConfig()

		force, _ := cmd.Flags().GetBool("force")
		if !force {
			fmt.Println(getStopUsage())
			fmt.Println("\nIf you wish to proceed, run again with the -f flag.")
			os.Exit(0)
		}
		if err := start.Stop(); err != nil {
			util.ErrAndExit("Exiting due to stop failure: %v", err)
		}
	},
}

func init() {
	stopCmd.Flags().BoolP("force", "f", false, "Use this flag when you're sure you want to run scope stop")

	RootCmd.AddCommand(stopCmd)
}
