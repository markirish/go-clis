package list

import (
	. "github.com/markirish/go-clis/cmd/clis/cobra/cmd/list/pods"
	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListCmd(globalOptions *GlobalCommandOptions) *cobra.Command {
	listOpts := &ListCommandOptions{}
	listOpts.GlobalCommandOptions = *globalOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List resources",
		Long: `List will display a list of resources based on the specified type and
filters. You can use this command to view all resources or filter them based
on specific criteria.`,
		RunE: func(cmd *cobra.Command, args []string) error {
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

	return cmd
}
