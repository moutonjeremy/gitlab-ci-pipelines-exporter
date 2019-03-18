dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

sec:
	go get github.com/securego/gosec/cmd/gosec/...
	gosec . -vendor

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/gitlab-ci-pipelines-exporter_linux_amd64
