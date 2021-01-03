package stats

import "time"

type NoopKstats struct {
}

func NewNoop() (Stats, func(), error) {
	return &NoopKstats{}, nil, nil
}
func (st *NoopKstats) Clone() (Stats, error) {

	return st, nil
}

func (st *NoopKstats) IsClientNil() bool {
	return false
}

func (st *NoopKstats) Count(stat string, value int64) error {
	return nil
}

func (st *NoopKstats) Gauge(stat string, value int64) error {
	return nil
}

func (st *NoopKstats) Increment(stat string, value int64) error {

	return nil
}

func (st *NoopKstats) Timing(stat string, delta int64) error {

	return nil
}

func (st *NoopKstats) TimingDuration(stat string, delta time.Duration) error {

	return nil
}
