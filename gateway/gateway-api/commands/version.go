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
		Short: "Return gateway-api version",
		Long:  "Return gateway-api version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("gateway-api\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, SigmaSee Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
