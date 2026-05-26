module github.com/markirish/go-clis/clis/flag

require github.com/markirish/go-clis/internal/app v0.0.0
replace github.com/markirish/go-clis/internal/app => ../../../internal/app

require github.com/markirish/go-clis/internal/pkg v0.0.0
replace github.com/markirish/go-clis/internal/pkg => ../../../internal/pkg

require github.com/markirish/go-clis/clis/flag/flagext v0.0.0
replace github.com/markirish/go-clis/clis/flag/flagext => ./flag-ext


go 1.26.2
