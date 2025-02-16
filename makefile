test:
	go test -race .

cover:
	go test -race -cover -coverprofile=out.html
	go tool cover -html=out.html
