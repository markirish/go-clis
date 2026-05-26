package options

// `list` command options

import (
	. "github.com/markirish/go-clis/internal/pkg/pods"
)

type ListCommandOptions struct {
	Filter string
	SortBy string
	Watch  bool
	Wide   bool
}

type ListPodsCommandOptions struct {
	Phase               string
	Node                string
	RestartsGreaterThan uint
}

func ListPodOptionsGenerator(globalOptions GlobalOptions, listOpts ListCommandOptions) ListPodOptions {
	return ListPodOptions{}
}

type ListDeploymentsCommandOptions struct {
	Unavailable bool
	MinReplicas uint
}

type ListServicesCommandOptions struct {
	Type string
}

type ListNodesCommandOptions struct {
	Node     string
	NotReady bool
}

type ListEventsCommandOptions struct {
	Type   string
	Reason string
}
