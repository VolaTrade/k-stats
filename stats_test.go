package stats_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	stats "github.com/volatrade/k-stats"
)

var (
	cfg *stats.Config
)

func TestMain(m *testing.M) {
	cfg = &stats.Config{Host: "localhost", Port: 8125, Env: "Dev"}
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGauge(t *testing.T) {
	st, err := stats.New(cfg)
	assert.Nil(t, err)

	err1 := st.Gauge("gauge.testing", 3)
	assert.Nil(t, err1)
}

func TestIncrement(t *testing.T) {
	st, err := stats.New(cfg)
	assert.Nil(t, err)
	err1 := st.Increment("increment.testing", 1)
	assert.Nil(t, err1)
}
