package stats

import (
	"fmt"

	"github.com/cactus/go-statsd-client/v4/statsd"
)

type (
	Config struct {
		Host string
		Port int
		Env  string
	}

	Stats struct {
		client statsd.Statter
		cfg    *Config
	}
)

func New(cfg *Config) (*Stats, error) {
	conf := &statsd.ClientConfig{
		Address: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Prefix:  cfg.Env,
	}
	println("creating stats connection to ->", conf.Address)
	client, err := statsd.NewClientWithConfig(conf)
	if err != nil {
		return nil, err
	}
	return &Stats{client: client, cfg: cfg}, nil
}

func (st *Stats) Increment(name string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.Inc(name, value, 1.0)
}

func (st *Stats) Gauge(name string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.Gauge(name, value, 1.0)
}
