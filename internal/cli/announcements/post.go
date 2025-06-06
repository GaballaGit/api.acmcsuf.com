package announcements

import(
	"github.com/spf13/cobra"
)


var PostAnnouncement = &cobra.Command{
	Use: "post",
	Short: "Post a new annoucement",
}
