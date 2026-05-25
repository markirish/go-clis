package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/markirish/go-clis/clis/flag/flag-ext"
	. "github.com/markirish/go-clis/internal/app/options"
)

func main() {

	// TODO: For now, assume that the user is passing in a verb + noun command
	// In the future, test more-complex command structures
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <command> <subcommand> [flags]\n", os.Args[0])
	}

	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Bind global options to a struct to pass options to the next layers
	globalOptions := &GlobalOptions{}
	BindGlobalFlags(flagSet, globalOptions)

	flagSet.Parse(os.Args[1:])

	// // Grabbing the command

	// if isFlag(os.Args[1]) {
	// 	fmt.Printf("Error: First argument must be a command, not a flag.\n")
	// 	os.Exit(1)
	// }
	// cmd := os.Args[1]

	// verbOpts := verbOptionsFactories[cmd]()
	// bindVerbFlags(cmd, verbOpts, flagSet)

	// // Bind options for the specific command to pass to the next layers

	// if isFlag(os.Args[2]) {
	// 	fmt.Printf("Error: Second argument must be a subcommand, not a flag.\n")
	// 	fmt.Printf("Usage: %s %s <subcommand> [flags]\n", os.Args[0], cmd)
	// 	os.Exit(1)
	// }

	// // After parsing noun
	// nounOpts := nounOptionsFactories[cmd][noun]()
	// bindNounFlags(cmd, noun, nounOpts, flagSet)

	// // Now parse all flags
	// flagSet.Parse(os.Args[3:])
}

// // Bind verb-level flags
// func bindVerbFlags(verb string, opts interface{}, flagSet *flag.FlagSet) error {
// 	flagDefs, ok := verbFlags[verb]
// 	if !ok {
// 		return nil
// 	}

// 	for _, definition := range flagDefs {
// 		if err := definition.FlagBinder.BindFlag(definition, flagSet, opts); err != nil {
// 			return fmt.Errorf("Error binding flag --%s: %v", definition.Long, err)
// 		}
// 	}
// 	return nil
// }

// // Bind noun-level flags
// func bindNounFlags(verb, noun string, opts interface{}, flagSet *flag.FlagSet) error {
// 	if _, ok := nounFlags[verb]; !ok {
// 		return nil
// 	}

// 	flagDefs, ok := nounFlags[verb][noun]
// 	if !ok {
// 		return nil
// 	}

// 	for _, definition := range flagDefs {
// 		if err := definition.FlagBinder.BindFlag(definition, flagSet, opts); err != nil {
// 			return fmt.Errorf("Error binding flag --%s: %v", definition.Long, err)
// 		}
// 	}
// 	return nil
// }
