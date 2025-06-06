package events

import(

	"github.com/spf13/cobra"
)


var PostEvent = &cobra.Command{
	Use: "post",
	Short: "Post a new event",
}
