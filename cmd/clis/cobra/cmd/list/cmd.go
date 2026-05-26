package list

import (
	"fmt"

	. "github.com/markirish/go-clis/cmd/clis/cobra/cmd/list/pods"
	. "github.com/markirish/go-clis/cmd/clis/cobra/cmd/list/services"
	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListCmd(globalOptions *GlobalOptions) *cobra.Command {
	listOpts := &ListCommandOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List resources",
		Long: `List will display a list of resources based on the specified type and
filters. You can use this command to view all resources or filter them based
on specific criteria.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Running list with global options: %+v\n", globalOptions)
			fmt.Printf("Running list with list options: %+v\n", listOpts)
			fmt.Printf("Args: %v\n", args)
			return nil
		},
	}

	flags := cmd.PersistentFlags()

	flags.StringVar(&listOpts.Filter, "filter", "", "Filter output")
	flags.StringVar(&listOpts.SortBy, "sort-by", "", "Sort output by specified field")
	flags.BoolVarP(&listOpts.Watch, "watch", "w", false, "Watch for changes")
	flags.BoolVar(&listOpts.Wide, "wide", false, "Display wide output")

	// Register subcommands
	cmd.AddCommand(NewListPodsCmd(globalOptions, listOpts))
	cmd.AddCommand(NewListServicesCmd(globalOptions, listOpts))

	return cmd
}
