package resolver

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"

	_ "golang.org/x/image/bmp"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

func TestExtensionResolver_Resolve(t *testing.T) {
	type args struct {
		directory string
		filename  string
	}
	tests := []struct {
		name    string
		r       ExtensionResolver
		args    args
		wantErr bool
	}{
		{"Will resolved", ExtensionResolver{}, args{"./", "test.jpg_large"}, false},
		{"Nothing to do", ExtensionResolver{}, args{"./", "test.jpg"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			path := filepath.Join(tt.args.directory, tt.args.filename)
			createImageFile(path)

			r := ExtensionResolver{}
			if err := r.Resolve(tt.args.directory, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("ExtensionResolver.Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
			resolvedFilepath := filepath.Join(tt.args.directory, strings.Replace(tt.args.filename, "_large", "", -1))
			_, e := os.Stat(resolvedFilepath)
			if e != nil {
				t.Errorf("%v not exists.", resolvedFilepath)
			}

			// teardown
			os.Remove(filepath.Join(tt.args.directory, tt.args.filename))
			os.Remove(resolvedFilepath)
		})
	}
}

func TestExtensionResolver_Available(t *testing.T) {
	type args struct {
		directory string
		filename  string
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		wantErr  bool
		setup    func(string)
		teardown func(string)
	}{
		{
			"Will resolved", args{"./", "test.jpg_large"}, true, false,
			func(path string) {
				createImageFile(path)
			},
			func(path string) {
				os.Remove(path)
			},
		},
		{
			"Nothing to do", args{"./", "test.jpg"}, false, false,
			func(path string) {
				createImageFile(path)
			},
			func(path string) {
				os.Remove(path)
			},
		},
		{
			"Could not open that file is not image", args{"./", "test.txt"}, false, true,
			func(path string) {
				file, e := os.Create(path)
				defer file.Close()
				if e != nil {
					panic(e)
				}
				file.WriteString("test")
			},
			func(path string) {
				os.Remove(path)
			},
		},
	}
	for _, tt := range tests {
		// setup
		tt.setup(filepath.Join(tt.args.directory, tt.args.filename))

		t.Run(tt.name, func(t *testing.T) {
			r := ExtensionResolver{}
			got, err := r.Available(tt.args.directory, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtensionResolver.Available() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtensionResolver.Available() = %v, want %v", got, tt.want)
			}
		})

		// teardown
		tt.teardown(filepath.Join(tt.args.directory, tt.args.filename))
	}
}

func createImageFile(path string) {
	rgba := image.NewRGBA(image.Rect(0, 0, 100, 100))
	file, e := os.Create(path)
	defer file.Close()
	if e != nil {
		panic(e)
	}

	e = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
	if e != nil {
		panic(e)
	}
}
