package resolver

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFilepathResolver_Resolve(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	type args struct {
		path string
	}
	tests := []struct {
		name string
		r    FilepathResolver
		args args
		want string
	}{
		{"contains home directory alias", FilepathResolver{}, args{"~/Download/test"}, "/Users/seinflaw/Download/test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FilepathResolver{}
			if got := r.Resolve(tt.args.path); got != tt.want {
				t.Errorf("FilepathResolver.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
