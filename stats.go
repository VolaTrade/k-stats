package stats

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cactus/go-statsd-client/v4/statsd"
	logger "github.com/volatrade/currie-logs"
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
		Count(stat string, value int64)
		Gauge(stat string, value int64)
		Increment(stat string, value int64)
		IsClientNil() bool
		Timing(stat string, delta int64)
		TimingDuration(stat string, delta time.Duration)
	}

	kstats struct {
		client statsd.Statter
		cfg    *Config
		logger *logger.Logger
	}
)

func New(cfg *Config, logger *logger.Logger) (Stats, func(), error) {

	if cfg.Env == "DEV" {

		return NewNoop()
	}

	conf := &statsd.ClientConfig{
		Address: strings.ToLower(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
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

	return &kstats{client: client, cfg: cfg, logger: logger}, end, nil
}

func (st *kstats) Clone() (Stats, error) {

	clone, _, err := New(st.cfg, st.logger)

	if err != nil {
		return nil, err
	}

	return clone, nil
}

func (st *kstats) IsClientNil() bool {
	return st.client == nil
}

func (st *kstats) Count(stat string, value int64) {

	if err := st.client.Inc(stat, value, 1.0); err != nil {
		st.logger.Errorw("Could not aggregate stats", "error", err.Error())
	}
}

func (st *kstats) Gauge(stat string, value int64) {

	if err := st.client.Gauge(stat, value, 1.0); err != nil {
		st.logger.Errorw("Could not aggregate stats", "error", err.Error())
	}
}

func (st *kstats) Increment(stat string, value int64) {

	if err := st.client.Inc(stat, value, 1.0); err != nil {
		st.logger.Errorw("Could not aggregate stats", "error", err.Error())
	}
}

func (st *kstats) Timing(stat string, delta int64) {

	if err := st.client.Timing(stat, delta, 1.0); err != nil {
		st.logger.Errorw("Could not aggregate stats", "error", err.Error())
	}
}

func (st *kstats) TimingDuration(stat string, delta time.Duration) {

	if err := st.client.TimingDuration(stat, delta, 1.0); err != nil {
		st.logger.Errorw("Could not aggregate stats", "error", err.Error())
	}
}
