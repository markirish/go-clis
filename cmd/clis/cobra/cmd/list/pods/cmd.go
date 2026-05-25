// /*
// Copyright © 2026 NAME HERE <EMAIL ADDRESS>
// */
// package pods

// import (
// 	"fmt"

// 	. "github.com/markirish/go-clis/clis/cobra/cmd/list"
// 	. "github.com/markirish/go-clis/internal/app/options"
// 	"github.com/spf13/cobra"
// )

// var listPodsOptions ListPodsCommandOptions

// // ListPodsCmd represents the pods command
// var ListPodsCmd = &cobra.Command{
// 	Use:   "pods",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: run,
// }

// func init() {
// 	ListCmd.AddCommand(ListPodsCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	ListPodsCmd.PersistentFlags().StringVar(&listPodsOptions.Phase, "phase", "", "Filter by pod phase")
// 	ListPodsCmd.PersistentFlags().StringVar(&listPodsOptions.Node, "node", "", "Filter by node name")
// 	ListPodsCmd.PersistentFlags().UintVar(&listPodsOptions.RestartsGreaterThan, "restarts-greater-than", 0, "Filter pods with restarts greater than the specified number")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// func run(_ *cobra.Command, args []string) {
// 	fmt.Printf("Running list command with listpodoptions: %+v and args: %v\n", listPodsOptions, args)
// 	fmt.Printf("Running list command with listoptions: %+v and args: %v\n", ListOptions, args)
// 	fmt.Printf("ListOptions: wide: %v, watch: %v\n", *ListOptions.Wide, *ListOptions.Watch)
// }

package pods

import (
	"fmt"

	"github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewPodsCmd(
	globalOptions *options.GlobalOptions,
	listOptions *options.ListCommandOptions,
) *cobra.Command {
	podsOpts := &options.ListPodsCommandOptions{}

	cmd := &cobra.Command{
		Use:   "pods",
		Short: "List pods",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("List options: %+v\n", listOptions)
			fmt.Printf("Pods options: %+v\n", podsOpts)
			fmt.Printf("Args: %v\n", args)

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&podsOpts.Phase, "phase", "", "Filter by pod phase")
	flags.StringVar(&podsOpts.Node, "node", "", "Filter by node name")
	flags.UintVar(&podsOpts.RestartsGreaterThan, "restarts-greater-than", 0, "Filter pods with restarts greater than the specified number")

	return cmd
}
