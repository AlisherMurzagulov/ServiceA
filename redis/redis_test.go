package redis

import (
	"serviceB/cfg"
	"testing"
)

func TestSetupRedis(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: testing.CoverMode()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Setup("../config.json")
			NewManager()
		})
	}
}
