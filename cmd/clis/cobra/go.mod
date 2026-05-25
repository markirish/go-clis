module github.com/markirish/go-clis/clis/cobra

go 1.26.2

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
)

// require github.com/markirish/go-clis/internal/clis/cobra/cmd v0.0.0

require github.com/markirish/go-clis/internal/app v0.0.0
replace github.com/markirish/go-clis/internal/app => ../../app
