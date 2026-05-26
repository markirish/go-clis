package options

import (
	. "github.com/markirish/go-clis/internal/pkg/pods"
)

// `list` command options

type ListCommandOptions struct {
	GlobalCommandOptions
	Filter string
	SortBy string
	Watch  bool
	Wide   bool
}

// `list pods` command options

type ListPodsCommandOptions struct {
	ListCommandOptions
	Phase               string
	Node                string
	RestartsGreaterThan uint
}

func ListPodOptionsGenerator(options ListPodsCommandOptions) ListPodOptions {
	return ListPodOptions{
		Namespace:           options.Namespace,
		Timeout:             options.Timeout,
		GlobalTimeout:       options.GlobalTimeout,
		Verbose:             options.Verbose,
		Filter:              options.Filter,
		SortBy:              options.SortBy,
		Phase:               options.Phase,
		Node:                options.Node,
		RestartsGreaterThan: options.RestartsGreaterThan,
	}
}

// `list deployments` command options

type ListDeploymentsCommandOptions struct {
	ListCommandOptions
	Unavailable bool
	MinReplicas uint
}

// `list services` command options

type ListServicesCommandOptions struct {
	ListCommandOptions
	Type string
}

// `list nodes` command options

type ListNodesCommandOptions struct {
	ListCommandOptions
	Node     string
	NotReady bool
}

// `list events` command options

type ListEventsCommandOptions struct {
	ListCommandOptions
	Type   string
	Reason string
}
