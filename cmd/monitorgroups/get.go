package monitorgroups

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ndbeals/uptimectl/pkg/betteruptime"
	"github.com/ndbeals/uptimectl/pkg/table"
)

const NoHeaderKey = "no-header"

var noHeader bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get a list of monitor groups",
	Aliases: []string{"g", "list", "ls"},
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()
		monitorGroups, err := client.ListMonitoringGroups()
		if err != nil {
			return err
		}
		body := make([][]string, 0, len(monitorGroups))
		for _, item := range monitorGroups {
			body = append(body, []string{
				item.Id,
				item.Attributes.Name,
				fmt.Sprint(item.Attributes.Paused),
				item.Attributes.CreatedAt.Local().String(),
			})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Paused", "Created at"}, body)
		}
		return nil
	},
}

func init() {
	MonitorGroupsCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
}
