/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	. "github.com/markirish/go-clis/clis/cobra/cmd"
	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

var ListOptions ListCommandOptions = ListCommandOptions{
	Filter: new(string),
	SortBy: new(string),
	Watch:  new(bool),
	Wide:   new(bool),
}

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
	Long: `List will display a list of resources based on the specified type and
filters. You can use this command to view all resources or filter them based
on specific criteria.`,
	Run: runList,
}

func init() {
	RootCmd.AddCommand(ListCmd)

	flags := ListCmd.PersistentFlags()

	flags.StringVar(ListOptions.Filter, "filter", "", "Filter output")
	flags.StringVar(ListOptions.SortBy, "sort-by", "", "Sort output by specified field")
	flags.BoolVarP(ListOptions.Watch, "watch", "w", false, "Watch for changes")
	flags.BoolVar(ListOptions.Wide, "wide", false, "Display wide output")
}

func runList(cmd *cobra.Command, args []string) {

	fmt.Printf("Running list command with options: %+v and args: %v\n", ListOptions, args)
}
