package announcements

import(
	"fmt"

	"github.com/spf13/cobra"
)

var GetAnnouncement = &cobra.Command{
	Use: "get",
	Short: "Get info on an announcement",

	Run: getAnnouncement,
}

func getAnnouncement(cmd *cobra.Command, args []string){
	fmt.Println("Get announcement called")
}

