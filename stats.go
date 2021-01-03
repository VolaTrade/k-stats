package stats

import (
	"fmt"
	"log"
	"strings"
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
		client statsd.Statter
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

	return &Stats{client: client, cfg: cfg}, end, nil
}

func NewNoop() *Stats {
	return &Stats{
		client: nil,
		cfg:    &Config{Env: "DEV"},
	}
}

func Clone(st *Stats) (*Stats, error) {

	st, _, err := New(st.cfg)
	if err != nil {
		return &Stats{client: nil, cfg: &Config{Env: "DEV"}}, nil
	}

	return st, nil
}

func (st *Stats) IsClientNil() bool {
	return st.client == nil
}

func (st *Stats) Count(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.Inc(stat, value, 1.0)
}

func (st *Stats) Gauge(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.Gauge(stat, value, 1.0)
}

func (st *Stats) Increment(stat string, value int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.Inc(stat, value, 1.0)
}

func (st *Stats) Timing(stat string, delta int64) error {
	if st.cfg.Env == "DEV" {
		return nil
	}

	return st.client.Timing(stat, delta, 1.0)
}

func (st *Stats) TimingDuration(stat string, delta time.Duration) error {
	if st.cfg.Env == "DEV" {
		return nil
	}
	return st.client.TimingDuration(stat, delta, 1.0)
}
