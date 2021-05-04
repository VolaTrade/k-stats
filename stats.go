package stats

import (
	"fmt"
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
		client statsd.StatSender
		cfg    *Config
		logger *logger.Logger
	}
)

func (mt *kstats) connect() error {
	var err error

	if mt.cfg.Env == "" {
		return fmt.Errorf("EnvService property must not be empty")
	}

	statsdAddress := fmt.Sprintf("%s:%d", mt.cfg.Host, mt.cfg.Port)

	mt.logger.Infow("STATSD - initializing statsd connection", "address", statsdAddress)

	mt.client, err = statsd.NewClient(statsdAddress, "test")

	if err != nil {
		/* If nothing is listening on the target port, an error is returned and
		the returned client does nothing but is still usable. So we can just log
		the error and go on. */
		mt.logger.Errorw("STATSD - failure when initializing statsd client", "error", err.Error())
		return err
	}
	return nil
}

func New(cfg *Config, logger *logger.Logger) (Stats, error) {
	stz := &kstats{cfg: cfg, logger: logger}

	if err := stz.connect(); err != nil {
		return nil, err

	}

	return stz, nil
}

func (st *kstats) Clone() (Stats, error) {

	clone, err := New(st.cfg, st.logger)

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
