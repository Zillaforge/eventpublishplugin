package cmd

import (
	"eventpublishplugin/server"

	"github.com/spf13/cobra"
)

/*
NewServeCmd starts the EventPublishPlugin and listens for incoming requests.
*/
func NewServeCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Start EventPublishPlugin",
		Long:  "Describe about starting EventPublishPlugin",
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
	}
	return
}
