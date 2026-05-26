package nodes

import (
	"fmt"

	"github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListNodesCmd(
	globalOptions *options.GlobalOptions,
	listOptions *options.ListCommandOptions,
) *cobra.Command {
	listNodesOpts := &options.ListNodesCommandOptions{}

	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "List nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("List options: %+v\n", listOptions)
			fmt.Printf("Nodes options: %+v\n", listNodesOpts)
			fmt.Printf("Args: %v\n", args)

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&listNodesOpts.Node, "node", "", "Filter by node name")
	flags.BoolVar(&listNodesOpts.NotReady, "not-ready", false, "Include only nodes that are not ready")

	return cmd
}
