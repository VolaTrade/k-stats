package stats

import (
	"fmt"

	"github.com/google/wire"

	"github.com/cactus/go-statsd-client/v4/statsd"
)

var Module = wire.NewSet(
	New,
)

type (
	Stats interface {
		Increment(name string, value int64) error
		Gauge(name string, value int64) error
	}

	Config struct {
		Host string
		Port int
		Env  string
	}

	VolaClient struct {
		client statsd.Statter
		cfg    *Config
	}
)

func New(cfg *Config) (*VolaClient, error) {
	conf := &statsd.ClientConfig{
		Address: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Prefix:  cfg.Env,
	}
	println("creating stats connection to ->", conf.Address)
	client, err := statsd.NewClientWithConfig(conf)
	if err != nil {
		return nil, err
	}
	return &VolaClient{client: client, cfg: cfg}, nil
}

func (vc *VolaClient) Increment(name string, value int64) error {
	if vc.cfg.Env == "DEV" {
		return nil
	}
	return vc.client.Inc(name, value, 1.0)
}

func (vc *VolaClient) Gauge(name string, value int64) error {
	if vc.cfg.Env == "DEV" {
		return nil
	}
	return vc.client.Gauge(name, value, 1.0)
}
