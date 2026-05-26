package deployments

import (
	"fmt"

	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/spf13/cobra"
)

func NewListDeploymentsCmd(
	globalOptions *GlobalOptions,
	listOptions *ListCommandOptions,
) *cobra.Command {
	listDeploymentsOptions := ListDeploymentsCommandOptions{}

	cmd := &cobra.Command{
		Use:   "deployments",
		Short: "List deployments",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("List options: %+v\n", listOptions)
			fmt.Printf("Deployments options: %+v\n", listDeploymentsOptions)
			fmt.Printf("Args: %v\n", args)

			return nil
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&listDeploymentsOptions.Unavailable, "name", false, "Filter by deployment name")
	flags.UintVar(&listDeploymentsOptions.MinReplicas, "namespace", 0, "Filter by namespace")

	return cmd
}
