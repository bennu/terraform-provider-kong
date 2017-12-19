package kong

import (
	"github.com/bennu/kong-go/client"
)

type Config struct {
	Endpoint string
}

func (c *Config) Client() (*client.Client, error) {
	cli, err := client.NewClient(nil)
	if err != nil {
		panic(err)

	}

	return cli, nil
}
