build:
	go build app.go
run:
	go run app.go
compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm app.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 app.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 app.go
test:
	go test ./...
test-conver:
	go test -cover ./...
update-data:
	./scripts/get_data/file.sh
	perl ./scripts/get_data/download.pl data.csv