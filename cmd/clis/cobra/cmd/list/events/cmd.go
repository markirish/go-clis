package events

import (
	"fmt"

	"github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListEventsCmd(
	globalOptions *options.GlobalCommandOptions,
	listOptions *options.ListCommandOptions,
) *cobra.Command {
	listEventsOpts := &options.ListEventsCommandOptions{}

	cmd := &cobra.Command{
		Use:   "events",
		Short: "List events",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("List options: %+v\n", listOptions)
			fmt.Printf("Events options: %+v\n", listEventsOpts)
			fmt.Printf("Args: %v\n", args)

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&listEventsOpts.Type, "type", "", "Filter by event type")
	flags.StringVar(&listEventsOpts.Reason, "reason", "", "Filter by event reason")

	return cmd
}
