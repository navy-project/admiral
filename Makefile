default: dev

bin:
	@sh -c "$(CURDIR)/scripts/build.sh"

dev:
	@TF_DEV=1 sh -c "$(CURDIR)/scripts/build.sh"

updatedeps:
	go get -d -v -p 2 ./...

.PHONY: bin default test updatedeps
