.PHONY: sort-import
sort-import:
	@./scripts/goimports-reviser.sh

.PHONY: test
test:
	@go test -v $$(go list ./... | grep -v vendor)
