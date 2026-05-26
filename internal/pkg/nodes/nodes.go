package nodes

type ListNodesOptions struct {
	Namespace     string
	Timeout       uint
	GlobalTimeout uint
	Verbose       bool
	Filter        string
	SortBy        string
	Node          string
	NotReady      bool
	// These fields should be handled at a higher level between the CLI and the
	// internal package, as they are more about output formatting and behavior
	// rather than pod-specific options
	//
	// Watch  bool
	// Wide   bool
}

func ListNodes(options ListNodesOptions) {
	// Placeholder for actual implementation

}
