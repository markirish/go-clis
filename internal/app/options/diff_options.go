package options

// `diff` command options
type DiffCommandOptions struct {
	GlobalOptions  GlobalOptions
	File           string
	Name           string
	IgnoreMetaData bool
}

// `diff pods` command options
type DiffPodsCommandOptions struct {
	DiffCommandOptions DiffCommandOptions
	IgnoreImageTag     bool
	Container          string
}

// `diff deployments` command options
type DiffDeploymentsCommandOptions struct {
	DiffCommandOptions DiffCommandOptions
	IgnoreReplicas     bool
	IgnoreStrategy     bool
}

// `diff services` command options
type DiffServicesCommandOptions struct {
	DiffCommandOptions DiffCommandOptions
	IgnoreClusterIP    bool
	IgnoreNodePort     bool
}
