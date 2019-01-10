FROM golang:1.10 as builder

WORKDIR /go/src/gitlab.com/labbsio/gitlab-ci-pipelines-exporter
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN /go/bin/dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app

FROM scratch as release
COPY --from=builder /go/src/labbsio/gitlab-ci-pipelines-exporter .
CMD ["./app"]
