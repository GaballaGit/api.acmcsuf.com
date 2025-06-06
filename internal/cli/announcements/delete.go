package announcements

import(
	"github.com/spf13/cobra"
)


var DeleteAnnouncement = &cobra.Command{
	Use: "delete",
	Short: "Delete an announcement",
}
