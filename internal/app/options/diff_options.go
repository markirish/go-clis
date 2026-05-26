package options

// `diff` command options
type DiffCommandOptions struct {
	GlobalCommandOptions
	File           string
	Name           string
	IgnoreMetaData bool
}

// `diff pods` command options
type DiffPodsCommandOptions struct {
	DiffCommandOptions
	IgnoreImageTag bool
	Container      string
}

// `diff deployments` command options
type DiffDeploymentsCommandOptions struct {
	DiffCommandOptions
	IgnoreReplicas bool
	IgnoreStrategy bool
}

// `diff services` command options
type DiffServicesCommandOptions struct {
	DiffCommandOptions
	IgnoreClusterIP bool
	IgnoreNodePort  bool
}
