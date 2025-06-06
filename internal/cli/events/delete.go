package events

import(

	"github.com/spf13/cobra"
)

var DeleteEvent = &cobra.Command{
	Use: "delete",
	Short: "Delete an event",
}
