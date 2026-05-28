package app

import (
	"fmt"
	"io"
	"time"

	"github.com/markirish/go-clis/internal/app/options"
	corev1 "k8s.io/api/core/v1"
)

func PrintListPods(writer io.Writer, pods *corev1.PodList, opts options.ListPodsCommandOptions) error {
	switch opts.Output {
	case "", "table":
		printPodsTable(writer, pods, opts.Wide)
		return nil
	default:
		return fmt.Errorf("unsupported output format: %q", opts.Output)
	}
}

func printPodsTable(writer io.Writer, pods *corev1.PodList, wide bool) {
	if wide {
		fmt.Fprintf(writer, "%-32s %-52s %-12s %-12s %s\n",
			"NAMESPACE",
			"NAME",
			"PHASE",
			"AGE",
			"NODE",
		)

		for _, pod := range pods.Items {

			age := "<unknown>"

			if pod.Status.StartTime != nil {
				age = formatAge(time.Since(pod.Status.StartTime.Time))
			}

			fmt.Fprintf(writer, "%-32s %-52s %-12s %-12s %s\n",
				pod.Namespace,
				pod.Name,
				pod.Status.Phase,
				age,
				pod.Spec.NodeName,
			)
		}

		return
	}

	fmt.Fprintf(writer, "%-52s %-12s %-12s\n",
		"NAME",
		"PHASE",
		"AGE",
	)

	for _, pod := range pods.Items {

		age := "<unknown>"

		if pod.Status.StartTime != nil {
			age = formatAge(time.Since(pod.Status.StartTime.Time))
		}

		fmt.Fprintf(writer, "%-52s %-12s %-12s\n",
			pod.Name,
			pod.Status.Phase,
			age,
		)
	}
}

func formatAge(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}

	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}

	if d < 24*time.Hour {
		return fmt.Sprintf("%dh", int(d.Hours()))
	}

	return fmt.Sprintf("%dd", int(d.Hours()/24))
}
