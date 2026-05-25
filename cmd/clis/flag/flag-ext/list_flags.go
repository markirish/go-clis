package flagext

var ListFlags = []*FlagDefinition{
	{
		Long:            "filter",
		Short:           "",
		HelpText:        "Filter results",
		OptionFieldName: "Filter",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "sort-by",
		Short:           "",
		HelpText:        "Sort output by specified field",
		OptionFieldName: "SortBy",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "watch",
		Short:           "w",
		HelpText:        "Watch for changes",
		OptionFieldName: "Watch",
		FlagBinder:      BoolBinder,
	},
	{
		Long:            "wide",
		Short:           "",
		HelpText:        "Wide output",
		OptionFieldName: "Wide",
		FlagBinder:      BoolBinder,
	},
}

var ListPodFlags = []*FlagDefinition{
	//phase, node, restarts-greater-than
	{
		Long:            "phase",
		Short:           "",
		HelpText:        "Filter by pod phase",
		OptionFieldName: "Phase",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "node",
		Short:           "",
		HelpText:        "Filter by node name",
		OptionFieldName: "Node",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "restarts-greater-than",
		Short:           "",
		HelpText:        "Filter pods with restarts greater than the specified number",
		OptionFieldName: "RestartsGreaterThan",
		FlagBinder:      UintBinder,
	},
}
