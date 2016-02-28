.PHONY: test

test:
	@rm -f /tmp/ecuadorianid-coverprofile.out
	@go test -v -coverprofile=/tmp/ecuadorianid-coverprofile.out
	@go tool cover -func=/tmp/ecuadorianid-coverprofile.out
	@rm -f /tmp/ecuadorianid-coverprofile.out
