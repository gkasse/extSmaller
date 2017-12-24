package resolver

import (
	"os"
	"path/filepath"
	"strings"
	"image"
	logger "github.com/sirupsen/logrus"
)

const (
	replaceTarget = "_large"
)

type ExtensionResolver struct{}

func (r ExtensionResolver) Resolve(directory string, filename string) error {
	available, e := r.Available(directory, filename)
	if e != nil {
		logger.Error("Error occurred.", e)
		return e
	}
	if !available {
		logger.Info("This file is not image or not contains large in ext.")
		return nil
	}

	fullPath := filepath.Join(directory, filename)
	logger.Debug("Open file.", map[string]string{"filepath": fullPath})
	file, e := os.Open(fullPath)
	defer file.Close()
	if e != nil {
		return e
	}

	logger.Info(
		"This file is image and contains 'large' in ext, so change file extension.",
		map[string]string{"filename": filename})
	ext := filepath.Ext(filename)
	newExt := strings.Replace(ext, replaceTarget, "", -1)
	e = os.Rename(fullPath, filepath.Join(directory, strings.Replace(filename, ext, newExt, 1)))
	if e != nil {
		return e
	}
	return nil
}

func (r ExtensionResolver) Available(directory string, filename string) (bool, error) {
	isImage, e := r.isImage(directory, filename)
	if e != nil {
		logger.Error("Error occurred.", e)
		return false, e
	}
	if !isImage {
		logger.Debug("This file is not image.")
		return false, nil
	}
	if !strings.Contains(filepath.Ext(filename), replaceTarget) {
		logger.Debug("This file's extension is not contains 'large'.")
		return false, nil
	}
	return true, nil
}

func (r ExtensionResolver) isImage(directory string, filename string) (bool, error) {
	file, e := os.Open(filepath.Join(directory, filename))
	if e != nil {
		logger.Error("File could not open.", map[string]string{"filename": filename})
		return false, e
	}
	_, format, e := image.DecodeConfig(file)
	if e != nil {
		logger.Debug("Error occurred. Analyzed image is false.", e)
		return false, e
	}
	logger.Debug("Analyzed complete.", map[string]string{"File format": format})
	return true, nil
}
