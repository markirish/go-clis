package options

import "time"

type GlobalCommandOptions struct {
	Namespace     string
	Output        string
	Timeout       time.Duration
	GlobalTimeout time.Duration
	Verbose       bool
}
