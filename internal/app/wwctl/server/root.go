package server

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/warewulf/warewulf/internal/app/wwctl/completions"
	"github.com/warewulf/warewulf/internal/pkg/warewulfd"
	"github.com/warewulf/warewulf/internal/pkg/warewulfd/server"
)

var (
	baseCmd = &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "server [OPTIONS]",
		Short:                 "Start Warewulf server",
		RunE:                  CobraRunE,
		Args:                  cobra.NoArgs,
		ValidArgsFunction:     completions.None,
	}
)

func GetCommand() *cobra.Command {
	return baseCmd
}

func CobraRunE(cmd *cobra.Command, args []string) error {
	oldMask := syscall.Umask(000)
	defer syscall.Umask(oldMask)

	if err := warewulfd.DaemonInitLogging(); err != nil {
		return fmt.Errorf("failed to configure logging: %w", err)
	}
	if err := server.RunServer(); err != nil {
		return fmt.Errorf("failed to start Warewulf server: %w", err)
	}
	return nil
}
