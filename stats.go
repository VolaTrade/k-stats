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

	Stats interface {
		Clone() (Stats, error)
		Count(stat string, value int64) error
		Gauge(stat string, value int64) error
		Increment(stat string, value int64) error
		IsClientNil() bool
		Timing(stat string, delta int64) error
		TimingDuration(stat string, delta time.Duration) error
	}

	Kstats struct {
		client statsd.Statter
		cfg    *Config
	}
)

func New(cfg *Config) (Stats, func(), error) {

	if cfg.Env == "DEV" {

		return NewNoop()
	}

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

	return &Kstats{client: client, cfg: cfg}, end, nil
}

func (st *Kstats) Clone() (Stats, error) {

	clone, _, err := New(st.cfg)

	if err != nil {
		return nil, err
	}

	return clone, nil
}

func (st *Kstats) IsClientNil() bool {
	return st.client == nil
}

func (st *Kstats) Count(stat string, value int64) error {

	return st.client.Inc(stat, value, 1.0)
}

func (st *Kstats) Gauge(stat string, value int64) error {

	return st.client.Gauge(stat, value, 1.0)
}

func (st *Kstats) Increment(stat string, value int64) error {

	return st.client.Inc(stat, value, 1.0)
}

func (st *Kstats) Timing(stat string, delta int64) error {

	return st.client.Timing(stat, delta, 1.0)
}

func (st *Kstats) TimingDuration(stat string, delta time.Duration) error {

	return st.client.TimingDuration(stat, delta, 1.0)
}
