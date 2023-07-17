# for local build, setup environment variables from .env
.PHONY: ensure-mod lint install-tools test run-local lint-code

install-tools:
	@echo ">  Installing tools..."
	(cd /tmp && go get -u github.com/golangci/golangci-lint/cmd/golangci-lint)

lint:
	@echo ">  Checking codes..."
	golangci-lint run --out-format code-climate | jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'

test:
	@echo ">  Running tests..."
	go test -cover -race -v ./...

run-local: ensure-mod
	cd blockchain-service/cmd/blockchain && go run main.go

ensure-mod:
	go mod download
	go mod vendor

test-cover:
	@echo ">  Running tests..."
	go test ./... -v -coverpkg=./... -coverprofile tmp/cover.out.tmp
	cat tmp/cover.out.tmp | grep -vE "/mock_|storage/z_" > tmp/cover.out && rm tmp/cover.out.tmp
	go tool cover -func tmp/cover.out && go tool cover -html=tmp/cover.out -o tmp/cover.html

lint-code:
	staticcheck ./blockchain-service/logic