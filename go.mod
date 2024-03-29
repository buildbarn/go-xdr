module github.com/buildbarn/go-xdr

go 1.21

// Use the same version as in bb-storage.
replace mvdan.cc/gofumpt => mvdan.cc/gofumpt v0.3.0

require (
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10
	github.com/golang/mock v1.6.0
	github.com/gordonklaus/ineffassign v0.1.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/sync v0.6.0
	mvdan.cc/gofumpt v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/tools v0.1.9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
