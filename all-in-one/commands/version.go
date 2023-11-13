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
		Short: "Return all-in-one version",
		Long:  "Return all-in-one version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("all-in-one\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, SigmaSee Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
