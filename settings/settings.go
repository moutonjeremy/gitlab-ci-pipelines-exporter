package settings

import (
	"github.com/urfave/cli"
)

var (
	Gitlab struct {
		Url     string
		Token   string
		Refresh int64
		Owned   bool
	}
	Port int
)

// NewContext => set configuration from env vars or command parameters
func NewContext() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "gitlab_url, gu",
			Value:       "https://gitlab.com",
			Destination: &Gitlab.Url,
			EnvVar:      "GITLAB_URL",
			Usage:       "Gitlab Url",
		},
		cli.StringFlag{
			Name:        "gitlab_token, gt",
			Value:       "xxxxxxxxx-xxx",
			Destination: &Gitlab.Token,
			EnvVar:      "GITLAB_TOKEN",
			Usage:       "Gitlab Personnal Token",
		},
		cli.Int64Flag{
			Name:        "gitlab_refresh, gr",
			Value:       30,
			Destination: &Gitlab.Refresh,
			EnvVar:      "GITLAB_REFRESH",
			Usage:       "Refresh time Gitlab Pipelines status in sec",
		},
		cli.BoolFlag{
			Name:        "gitlab_owned, go",
			Destination: &Gitlab.Owned,
			EnvVar:      "GITLAB_OWNED",
			Usage:       "If you want just yours projects",
		},
		cli.IntFlag{
			Name:        "port, p",
			Value:       9999,
			Destination: &Port,
			EnvVar:      "PORT",
			Usage:       "Exporter port",
		},
	}
}
