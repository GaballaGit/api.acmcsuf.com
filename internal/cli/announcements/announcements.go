package announcements

import(
	"github.com/spf13/cobra"
)


var Announcements = &cobra.Command{
	Use: "announcements",
	Short: "Manage Announcements",
}


func init() {
	Announcements.AddCommand(GetAnnouncement)
}
