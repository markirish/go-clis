package options

// `describe` command options

type DescribeCommandOptions struct {
	GlobalCommandOptions
	Filter     string
	Name       string
	ShowEvents bool
}

type DescribePodsCommandOptions struct {
	DescribeCommandOptions
	ShowContainers bool
	ShowNodes      bool
}

type DescribeDeploymentsCommandOptions struct {
	DescribeCommandOptions
	ShowReplicas bool
	ShowPods     bool
}

type DescribeServicesCommandOptions struct {
	DescribeCommandOptions
	ShowEndpoints bool
	ShowSelectors bool
}

type DescribeNodesCommandOptions struct {
	DescribeCommandOptions
	ShowPods     bool
	ShowCapacity bool
}
