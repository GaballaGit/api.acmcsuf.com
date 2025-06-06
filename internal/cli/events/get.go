package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get [(-o|--output=)json|yaml|template|...]",
	Short: "Get an announcment",

	Run: getEvent,
}

func getEvent(cmd *cobra.Command, args []string) {
	fmt.Println("Get event called")
}
