package cfg

import "testing"

func TestSetup(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: testing.CoverMode(), args: struct{ path string }{path: "../config.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup(tt.args.path)
		})
	}
}
