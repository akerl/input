module github.com/akerl/input

go 1.14

// Needed until https://github.com/ktr0731/go-fuzzyfinder/pull/13 is merged
replace github.com/ktr0731/go-fuzzyfinder => github.com/akerl/go-fuzzyfinder v0.1.2-0.20200507171954-7f19dd52209e

require (
	github.com/akerl/timber/v2 v2.0.1
	github.com/dixonwille/wmenu/v5 v5.1.0
	github.com/ktr0731/go-fuzzyfinder v0.0.0-00010101000000-000000000000
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1
)
