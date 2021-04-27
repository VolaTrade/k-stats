package stats

import "time"

type noopKstats struct {
}

func NewNoop() (Stats, func(), error) {
	return &noopKstats{}, nil, nil
}
func (st *noopKstats) Clone() (Stats, error) {

	return st, nil
}

func (st *noopKstats) IsClientNil() bool {
	return false
}

func (st *noopKstats) Count(stat string, value int64) {
	return
}

func (st *noopKstats) Gauge(stat string, value int64) {
	return
}

func (st *noopKstats) Increment(stat string, value int64) {

	return
}

func (st *noopKstats) Timing(stat string, delta int64) {

	return
}

func (st *noopKstats) TimingDuration(stat string, delta time.Duration) {

	return
}
