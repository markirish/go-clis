package app

// `list` command options

type ListCommandOptions struct {
	GlobalOptions *GlobalOptions
	Filter        *string
	SortBy        *string
	Watch         *bool
	Wide          *bool
}

type ListPodsCommandOptions struct {
	ListCommandOptions  *ListCommandOptions
	Phase               *string
	Node                *string
	RestartsGreaterThan *uint
}

type ListDeploymentsCommandOptions struct {
	ListCommandOptions *ListCommandOptions
	Unavailable        *bool
	MinReplicas        *uint
}

type ListServicesCommandOptions struct {
	ListCommandOptions *ListCommandOptions
	Type               *string
}

type ListNodesCommandOptions struct {
	ListCommandOptions *ListCommandOptions
	Node               *string
	NotReady           *bool
}

type ListEventsCommandOptions struct {
	ListCommandOptions *ListCommandOptions
	Type               *string
	Reason             *string
}
