.PHONY: test, cover, fmt

test:
	go test -race .

cover:
	go test -race -cover -coverprofile=out.html
	go tool cover -html=out.html

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
