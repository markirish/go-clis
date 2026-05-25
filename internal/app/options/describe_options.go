package options

// `describe` command options

type DescribeCommandOptions struct {
	GlobalOptions GlobalOptions
	Filter        string
	Name          string
	ShowEvents    bool
}

type DescribePodsCommandOptions struct {
	DescribeCommandOptions DescribeCommandOptions
	ShowContainers         bool
	ShowNodes              bool
}

type DescribeDeploymentsCommandOptions struct {
	DescribeCommandOptions DescribeCommandOptions
	ShowReplicas           bool
	ShowPods               bool
}

type DesribeServicesCommandOptions struct {
	DescribeCommandOptions DescribeCommandOptions
	ShowEndpoints          bool
	ShowSelectors          bool
}

type DescribeNodesCommandOptions struct {
	DescribeCommandOptions DescribeCommandOptions
	ShowPods               bool
	ShowCapacity           bool
}
