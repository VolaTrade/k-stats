package stats

import (
	"fmt"
	"strings"
	"log"
	"time"

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

func New(cfg *Config) (*Stats, func(), error) {

	conf := &statsd.ClientConfig{
		Address: strings.ToLower(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
		Prefix:  strings.ToLower(fmt.Sprintf("%s.%s", cfg.Env, cfg.Service)),
	}
	log.Println("creating stats connection to ->", conf.Address)
	client, err := statsd.NewClientWithConfig(conf)
	if err != nil {
		return nil, nil, err
	}
	end := func() {

		if client != nil {
			if err := client.Close(); err != nil {
				log.Fatalf("Error closing new stats client: %v", err)
			}
			log.Println("K-stats client successful shutdown")
		}
	}

	return &Stats{Client: client, cfg: cfg}, end, nil
}

func NewNoop(cfg *Config) (*Stats, error) {
	return &Stats{Client: nil, cfg: &Config{Env: "DEV"}}, nil
}

func Clone(st *Stats) (*Stats, error) {
	conf := &statsd.ClientConfig{
		Address: fmt.Sprintf("%s:%d", st.cfg.Host, st.cfg.Port),
		Prefix:  fmt.Sprintf("%s.%s", st.cfg.Env,st. cfg.Service),
	}
	if st.cfg.Env == "DEV" {
		return nil, nil
	}

	log.Println("creating stats connection to ->", conf.Address)
	client, err := statsd.NewClientWithConfig(conf)
	if err != nil {
		return nil, err
	}

	return &Stats{Client: client, cfg: st.cfg}, nil
}

func (st *Stats) Count(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Inc(stat, value, 1.0)
}

func (st *Stats) Gauge(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Gauge(stat, value, 1.0)
}

func (st *Stats) Increment(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Inc(stat, value, 1.0)
}

func (st *Stats) Timing(stat string, delta int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.Timing(stat, delta, 1.0)
}

func (st *Stats) TimingDuration(stat string, delta time.Duration) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.Client.TimingDuration(stat, delta, 1.0)
}
