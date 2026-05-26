package pods

type ListPodOptions struct {
	Filter              string
	SortBy              string
	Phase               string
	Node                string
	RestartsGreaterThan uint
	// Watch  bool
	// Wide   bool
}

func ListPods(options ListPodOptions) {
	// Placeholder for actual implementation

}
