module github.com/akerl/input

go 1.14

// Needed until https://github.com/ktr0731/go-fuzzyfinder/pull/13 is merged
replace github.com/ktr0731/go-fuzzyfinder => github.com/akerl/go-fuzzyfinder v0.1.2-0.20200220111247-2e90b475f471

require (
	github.com/akerl/timber/v2 v2.0.1
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/dixonwille/wmenu v4.0.2+incompatible
	github.com/gdamore/tcell v1.3.0
	github.com/ktr0731/go-fuzzyfinder v0.0.0-00010101000000-000000000000
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/micmonay/keybd_event v1.1.0
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3
	gopkg.in/dixonwille/wlog.v2 v2.0.0 // indirect
)
