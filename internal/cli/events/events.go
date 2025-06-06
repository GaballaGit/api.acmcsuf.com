package events

import (
	"github.com/spf13/cobra"
)

var Events = &cobra.Command{
	Use:   "events",
	Short: "Manage Events",
}

func init() {
	Events.AddCommand(GetEvent)
}
