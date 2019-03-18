dep:
	go get -u github.com/golang/dep/cmd/dep
	/go/bin/dep ensure

sec:
	go get github.com/securego/gosec/cmd/gosec/...
	gosec . -vendor

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/gitlab-ci-pipelines-exporter_linux_amd64
