FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY bin/gitlab-ci-pipelines-exporter_linux_amd64 .
CMD ["./gitlab-ci-pipelines-exporter_linux_amd64"]
