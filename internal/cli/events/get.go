package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get an announcment",

	Run: getEvent,
}

func getEvent(cmd *cobra.Command, args []string) {
	fmt.Println("Get event called")
}
