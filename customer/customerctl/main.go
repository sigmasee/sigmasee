package main

import (
	"github.com/sigmasee/sigmasee/customer/customerctl/commands"
	"github.com/sigmasee/sigmasee/shared/enterprise/pkg/util"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
