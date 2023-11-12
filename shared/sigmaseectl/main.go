package main

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/pkg/util"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
