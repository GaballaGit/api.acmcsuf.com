package events

import(
	"fmt"
)

var getEvent = &cobra.Command{
	Use: "get",
	Short: "Get an announcment",

	Run: GetEvent,
}

func GetEvent(cmd *cobra.Command, args []string) {
	fmt.Println("Get event called")
}

