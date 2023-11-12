package main

import (
	"github.com/sigmasee/sigmasee/apex/apexctl/commands"
	"github.com/sigmasee/sigmasee/shared/enterprise/pkg/util"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
