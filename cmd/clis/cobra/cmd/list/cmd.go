/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
// package list

// import (
// 	"fmt"

// 	. "github.com/markirish/go-clis/clis/cobra/cmd"
// 	. "github.com/markirish/go-clis/internal/app/options"
// 	"github.com/spf13/cobra"
// )

// var ListOptions ListCommandOptions

// // ListCmd represents the list command
// var ListCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "List resources",
// 	Long: `List will display a list of resources based on the specified type and
// filters. You can use this command to view all resources or filter them based
// on specific criteria.`,
// 	Run: runList,
// }

// func init() {
// 	RootCmd.AddCommand(ListCmd)

// 	flags := ListCmd.PersistentFlags()

// 	flags.StringVar(&ListOptions.Filter, "filter", "", "Filter output")
// 	flags.StringVar(&ListOptions.SortBy, "sort-by", "", "Sort output by specified field")
// 	flags.BoolVarP(&ListOptions.Watch, "watch", "w", false, "Watch for changes")
// 	flags.BoolVar(&ListOptions.Wide, "wide", false, "Display wide output")
// }

// func runList(cmd *cobra.Command, args []string) {

//		fmt.Printf("Running list command with options: %+v and args: %v\n", ListOptions, args)
//	}
package list

import (
	"fmt"

	"github.com/markirish/go-clis/cmd/clis/cobra/cmd/list/pods"
	"github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListCmd(globalOptions *options.GlobalOptions) *cobra.Command {
	listOpts := &options.ListCommandOptions{}

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

	cmd.AddCommand(pods.NewPodsCmd(globalOptions, listOpts))

	return cmd
}
