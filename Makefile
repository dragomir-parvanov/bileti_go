coverage:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...

show-coverage:
	make coverage
	go tool cover -html=cover.out -o cover.html
	open cover.html