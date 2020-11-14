package warewulfd

import (
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
)

var (
	WarewulfdCmd = &cobra.Command{
		Use:                "warewulfd",
		Short:              "Warewulf Daemon Service",
		Long:               "This is the primary Warewulf service for provisioning nodes",
		RunE:				CobraRunE,
		PersistentPreRunE:  rootPersistentPreRunE,
	}
	verboseArg bool
	debugArg bool
)


func init() {
	WarewulfdCmd.PersistentFlags().BoolVarP(&verboseArg, "verbose", "v", false, "Run with increased verbosity.")
	WarewulfdCmd.PersistentFlags().BoolVarP(&debugArg, "debug", "d", false, "Run with debugging messages enabled.")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetRootCommand() *cobra.Command {
	return WarewulfdCmd
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	if verboseArg == true {
		wwlog.SetLevel(wwlog.VERBOSE)
	} else if debugArg == true {
		wwlog.SetLevel(wwlog.DEBUG)
	} else {
		wwlog.SetLevel(wwlog.INFO)
	}
	return nil
}