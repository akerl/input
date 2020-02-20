module github.com/akerl/input

go 1.13

// Needed until https://github.com/ktr0731/go-fuzzyfinder/pull/13 is merged
replace github.com/ktr0731/go-fuzzyfinder => github.com/akerl/go-fuzzyfinder v0.1.2-0.20200220111247-2e90b475f471

require (
	github.com/akerl/timber/v2 v2.0.1
	github.com/daviddengcn/go-colortext v0.0.0-20180409174941-186a3d44e920 // indirect
	github.com/dixonwille/wmenu v4.0.2+incompatible
	github.com/gdamore/tcell v1.3.0
	github.com/golangplus/bytes v0.0.0-20160111154220-45c989fe5450 // indirect
	github.com/golangplus/fmt v0.0.0-20150411045040-2a5d6d7d2995 // indirect
	github.com/golangplus/testing v0.0.0-20180327235837-af21d9c3145e // indirect
	github.com/ktr0731/go-fuzzyfinder v0.0.0-00010101000000-000000000000
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	gopkg.in/dixonwille/wlog.v2 v2.0.0 // indirect
)
