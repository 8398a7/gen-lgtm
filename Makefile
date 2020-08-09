.PHONY: deps
deps:
	go mod vendor
	go mod tidy

.PHONY: install
install: deps
	GOBIN=$(CURDIR)/bin go install ./...

.PHONY: gen-assets
gen-assets:
	go-assets-builder assets > assets.go
