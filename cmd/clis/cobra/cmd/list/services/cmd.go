package services

import (
	"fmt"

	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListServicesCmd(
	globalOptions *GlobalOptions,
	listOptions *ListCommandOptions,
) *cobra.Command {
	listServicesOpts := &ListServicesCommandOptions{}

	cmd := &cobra.Command{
		Use:   "services",
		Short: "List services",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("List options: %+v\n", listOptions)
			fmt.Printf("Services options: %+v\n", listServicesOpts)
			fmt.Printf("Args: %v\n", args)

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&listServicesOpts.Type, "type", "", "Filter by service type")

	return cmd
}
