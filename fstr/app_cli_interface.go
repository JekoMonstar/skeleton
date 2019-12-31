package fstr

var Cmd_Cli_Interface_Str = `package cli

import (
	ife "%s/app/interface"

	"%s/pub/glb"

	"github.com/spf13/cobra"
)

var InterfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		glb.InitConfig("interface", Cmd2)
		ife.Run()
	},
}

func init() {
	RootCmd.AddCommand(InterfaceCmd)
}
`
