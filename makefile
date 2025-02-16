.PHONY: test, cover, fmt, install-tools

test:
	go test -race .

cover:
	go test -race -cover -coverprofile=out.html
	go tool cover -html=out.html

install-tools:
	go install tool

fmt:
	go tool gofumpt -w .
	go tool gci write \
			--custom-order \
			--section standard \
			--section default \
			--section blank \
			--section dot \
			--skip-generated \
			.
