package resolver

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"testing"
)

func TestFilepathResolver_Resolve(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want func() string
	}{
		{"contains home directory alias", args{"~/Download/test"}, func() string {
			return filepath.Join(os.Getenv("HOME"), "Download/test")
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FilepathResolver{}
			if got := r.Resolve(tt.args.path); got != tt.want() {
				t.Errorf("FilepathResolver.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
