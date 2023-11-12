package commands

import (
	"fmt"
	"time"

	"github.com/sigmasee/sigmasee/shared/enterprise/pkg/util"
	"github.com/spf13/cobra"
)

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Return sigmaseectl version",
		Long:  "Return sigmaseectl version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("sigmaseectl\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, sigmasee Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
