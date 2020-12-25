package stats

import (
	"fmt"
	"strings"

	"github.com/cactus/go-statsd-client/v4/statsd"
)

type (
	Config struct {
		Env     string
		Host    string
		Port    int
		Service string
	}

	Stats struct {
		Client statsd.Statter
		cfg    *Config
	}
)

func New(cfg *Config) (*Stats, error) {
	conf := &statsd.ClientConfig{
		Address: strings.ToLower(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
		Prefix:  strings.ToLower(fmt.Sprintf("%s:%s", cfg.Env, cfg.Service)),
	}
	println("creating stats connection to ->", conf.Address)
	client, err := statsd.NewClientWithConfig(conf)
	if err != nil {
		return nil, err
	}
	return &Stats{Client: client, cfg: cfg}, nil
}

func Clone(st *Stats) (*Stats, error) {
	st, err := New(st.cfg)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (st *Stats) Increment(name string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Inc(name, value, 1.0)
}

func (st *Stats) Gauge(name string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Gauge(name, value, 1.0)
}
