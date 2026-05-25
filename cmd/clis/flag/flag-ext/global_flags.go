package flagext

import (
	"flag"
	"fmt"

	. "github.com/markirish/go-clis/internal/app/options"
)

var GlobalFlags = []*FlagDefinition{
	{
		Long:            "namespace",
		Short:           "n",
		HelpText:        "Namespace of the resource",
		OptionFieldName: "Namespace",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "output",
		Short:           "o",
		HelpText:        "Output format",
		OptionFieldName: "Output",
		FlagBinder:      StringBinder,
	},
	{
		Long:            "verbose",
		Short:           "v",
		HelpText:        "Enable verbose output",
		OptionFieldName: "Verbose",
		FlagBinder:      BoolBinder,
	},
}

func BindGlobalFlags(flagSet *flag.FlagSet, options *GlobalOptions) error {
	for _, definition := range GlobalFlags {
		if err := definition.FlagBinder.BindFlag(definition, flagSet, options); err != nil {
			return fmt.Errorf("Error binding flag --%s: %v\n", definition.Long, err)
		}
	}
	return nil
}
