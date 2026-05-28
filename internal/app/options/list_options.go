package options

import (
	"github.com/markirish/go-clis/internal/pkg/deployments"
	"github.com/markirish/go-clis/internal/pkg/nodes"
	"github.com/markirish/go-clis/internal/pkg/pods"
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

func ListPodOptionsGenerator(options ListPodsCommandOptions) pods.GetPodsOptions {
	return pods.GetPodsOptions{
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

func ListDeploymentsOptionsGenerator(options ListDeploymentsCommandOptions) deployments.GetDeploymentsOptions {
	return deployments.GetDeploymentsOptions{
		Namespace:     options.Namespace,
		Timeout:       options.Timeout,
		GlobalTimeout: options.GlobalTimeout,
		Verbose:       options.Verbose,
		Filter:        options.Filter,
		SortBy:        options.SortBy,
		Unavailable:   options.Unavailable,
		MinReplicas:   options.MinReplicas,
	}
}

// `list nodes` command options

type ListNodesCommandOptions struct {
	ListCommandOptions
	Node     string
	NotReady bool
}

func ListNodesOptionsGenerator(options ListNodesCommandOptions) nodes.GetNodesOptions {
	return nodes.GetNodesOptions{
		Namespace:     options.Namespace,
		Timeout:       options.Timeout,
		GlobalTimeout: options.GlobalTimeout,
		Verbose:       options.Verbose,
		Filter:        options.Filter,
		SortBy:        options.SortBy,
		Node:          options.Node,
		NotReady:      options.NotReady,
	}
}
