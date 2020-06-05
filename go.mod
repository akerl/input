module github.com/akerl/input

go 1.14

// Needed until https://github.com/ktr0731/go-fuzzyfinder/pull/13 is merged
replace github.com/ktr0731/go-fuzzyfinder => github.com/akerl/go-fuzzyfinder v0.1.2-0.20200507171954-7f19dd52209e

require (
	github.com/akerl/timber/v2 v2.0.1
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/dixonwille/wmenu v4.0.2+incompatible
	github.com/ktr0731/go-fuzzyfinder v0.0.0-00010101000000-000000000000
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1
	github.com/stretchr/testify v1.6.1 // indirect
	gopkg.in/dixonwille/wlog.v2 v2.0.0 // indirect
)
