package pods

import (
	"fmt"

	. "github.com/markirish/go-clis/internal/app/options"
	"github.com/markirish/go-clis/internal/pkg/pods"
	"github.com/spf13/cobra"
)

func NewListPodsCmd(
	globalOptions *GlobalCommandOptions,
	listOptions *ListCommandOptions,
) *cobra.Command {
	listPodsCommandOpts := &ListPodsCommandOptions{}

	cmd := &cobra.Command{
		Use:   "pods",
		Short: "List pods",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Now that we have all of the options bound, condense down into a single
			// struct so we only have to deal with ListPodsCommandOptions
			// Need to do the more specific options first, since they will overwrite
			// the global options if done second
			listPodsCommandOpts.ListCommandOptions = *listOptions
			listPodsCommandOpts.GlobalCommandOptions = *globalOptions

			fmt.Printf("Global options: %+v\n", globalOptions)
			fmt.Printf("Pods options: %+v\n", listPodsCommandOpts)

			listPodOptions := ListPodOptionsGenerator(*listPodsCommandOpts)

			err := pods.ListPods(cmd.Context(), listPodOptions)
			if err != nil {
				fmt.Println("TODO: Don't just print here, need to log, pass error to caller, etc. Fine for now though")
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&listPodsCommandOpts.Phase, "phase", "", "Filter by pod phase")
	flags.StringVar(&listPodsCommandOpts.Node, "node", "", "Filter by node name")
	flags.UintVar(&listPodsCommandOpts.RestartsGreaterThan, "restarts-greater-than", 0, "Filter pods with restarts greater than the specified number")

	return cmd
}
