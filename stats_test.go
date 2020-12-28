package stats_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	stats "github.com/volatrade/k-stats"
)

var (
	cfg *stats.Config
)

func TestMain(m *testing.M) {
	cfg = &stats.Config{Env: "Dev", Host: "localhost", Port: 8125, Service: "testservice"}
	retCode := m.Run()
	os.Exit(retCode)
}

func TestClone(t *testing.T) {
	st, end, err := stats.New(cfg)
	assert.Nil(t, err)

	stClone, end1, err1 := stats.Clone(st)
	assert.Nil(t, err1)
	
	if stClone.Client == nil {
		err = errors.New("failed to clone client")
	}
	assert.Nil(t, err)
	end()
	end1()
}

func TestCount(t *testing.T) {
	st, end, err := stats.New(cfg)
	assert.Nil(t, err)
	err1 := st.Count("count.testing", 4)
	assert.Nil(t, err1)
	end()
}

func TestGauge(t *testing.T) {
	st, end, err := stats.New(cfg)
	assert.Nil(t, err)

	err1 := st.Gauge("gauge.testing", 3)
	assert.Nil(t, err1)
	end()
}

func TestIncrement(t *testing.T) {
	st, end, err := stats.New(cfg)
	assert.Nil(t, err)
	err1 := st.Increment("increment.testing", 1)
	assert.Nil(t, err1)
	end()
}

func TestTiming(t *testing.T) {
	st, end, err := stats.New(cfg)
	assert.Nil(t, err)
	err1 := st.Timing("timing.testing", 2000)
	assert.Nil(t, err1)
	end()
}
