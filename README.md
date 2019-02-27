# gitlab-ci-pipelines-exporter
Export gitlab-ci pipeline status for prometheus (/metrics)

[![Build Status](https://travis-ci.org/Labbs/gitlab-ci-pipelines-exporter.svg?branch=master)](https://travis-ci.org/Labbs/gitlab-ci-pipelines-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/Labbs/gitlab-ci-pipelines-exporter)](https://goreportcard.com/report/github.com/Labbs/gitlab-ci-pipelines-exporter)
![Docker Pulls](https://img.shields.io/docker/pulls/labbsio/gitlab-ci-pipelines-exporter.svg)

## Options

| Name | Env | Parameter | Default | Description | 
|--|--|--|--|--|
| Gitlab Url | GITLAB_URL | gitlab_url, gu | https://gitlab.com | If you want use your own Gitlab instance |
| Gitlab Token | GITLAB_TOKEN | gitlab_token, gt | - | Create token in your profile with API and read options |
| Gitlab Refresh | GITLAB_REFRESH | gitlab_refresh, gr | 30 | In seconds, refresh every x seconds projects list and pipelines |
| Gitlab Owned | GITLAB_OWNED | gitlab_owned, go | false | If you want just yours projects |
| Port | PORT | port, p | 9999 | Exporter listening port |
