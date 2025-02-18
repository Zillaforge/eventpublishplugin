package cmd

import (
	cnt "eventpublishplugin/constants"
	"fmt"

	"github.com/spf13/cobra"
)

/*
NewVersionCmd shows the version of EventPublishPlugin.
*/
func NewVersionCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "version",
		Short: "Show Version",
		Long:  "Show ASUS Enterprise EventPublishPlugin Version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("EventPublishPlugin: %s\n", cnt.Version)
		},
	}
	return
}
